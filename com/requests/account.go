package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/isarq/nem-sdk-go/base"
	"github.com/isarq/nem-sdk-go/model"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// AccountInfo describes basic information for an account.
type AccountInfo struct {
	// Address contains the address of the account.
	Address string
	// Balance contains the balance of the account in micro NEM.
	Balance float64
	// vestedBalance contains the vested part of the balance of the account in micro NEM.
	VestedBalance float64
	// Importance contains the importance of the account.
	Importance float64
	// PublicKey contains the public key of the account.
	PublicKey string
	// Label has the label of the account( not used, always null).
	Label string
	// HarvestedBlocks contains the number of blocks that the account already harvested.
	HarvestedBlocks int
}

// AccountMetaData describes additional information for the account.
type AccountMetaData struct {
	// Status contains the harvesting status of a queried account.
	// The harvesting status can be one of the following values:
	// "UNKNOWN": The harvesting status of the account is not known.
	// "LOCKED": The account is not harvesting.
	// "UNLOCKED": The account is harvesting.
	Status string
	// RemoteStatus contains the status of teh remote harvesting of a queried account.
	// The remote harvesting status can be one of the following values:
	// "REMOTE": The account is a remote account and therefore remoteStatus is not applicable for it.
	// "ACTIVATING": The account has activated remote harvesting but it is not yet active.
	// "ACTIVE": The account has activated remote harvesting and remote harvesting is active.
	// "DEACTIVATING": The account has deactivated remote harvesting but remote harvesting is still active.
	// "INACTIVE": The account has inactive remote harvesting, or it has deactivated remote harvesting
	// and deactivation is operational.
	RemoteStatus string
	// CosignatoryOf is a JSON array of AccountInfo structures.
	// The account is cosignatory for each of the accounts in the array.
	CosignatoryOf []AccountInfo
	// Cosignatories is a JSON array of AccountInfo structures.
	// The array holds all accounts that are a cosignatory for this account.
	Cosignatories []AccountInfo
}

// AccountMetaDataPair includes durable information for an account and additional
// information about its state.
type AccountMetaDataPair struct {
	// Account contains the account object.
	Account AccountInfo `json:"account"`
	// Meta contain the account meta data object.
	Meta AccountMetaData `json:"meta"`
}

// HarvestInfo is information about harvested blocks
type HarvestInfo struct {
	TimeStamp  int64
	Difficulty int
	TotalFee   int
	ID         int
	Height     int64
}

// TransactionMetaData struct contains additional information about the transaction.
type TransactionMetaData struct {
	// The height of the block in which the transaction was included.
	Height int64 `json:"height"`
	// The id of the transaction.
	ID int `json:"id"`
	// The transaction hash.
	Hash struct {
		Data string `json:"data"`
	} `json:"hash"`
}

// Transactions meta data object contains additional information about the transaction.
type TransactionMetaDataPair struct {
	Meta        TransactionMetaData `json:"meta"`
	Transaction base.Transaction    `json:"transaction"`
}

// The unconfirmed transaction meta data contains the hash of the inner transaction in case the transaction
// is a multisig transaction. This data is need to initiate a multisig signature transaction.
type MetaData struct {
	Data string `json:"data"`
}

type unconfirmedMosaicTransactionMetaDataPair struct {
	Meta        TransactionMetaData    `json:"meta"`
	Transaction base.TransactionMosaic `json:"transaction"`
}

func (t *unconfirmedMosaicTransactionMetaDataPair) toStruct() (*TransactionMetaData, base.Transaction, error) {
	return &t.Meta, &t.Transaction, nil
}

type multiSignTransactionMetaDataPair struct {
	Meta        TransactionMetaData       `json:"meta"`
	Transaction base.MultiSignTransaction `json:"transaction"`
}

func (t *multiSignTransactionMetaDataPair) toStruct() (*TransactionMetaData, base.Transaction, error) {

	p := t.Transaction.OtherTrans
	d, err := json.Marshal(p)
	if err != nil {
		return nil, nil, err
	}

	var m json.RawMessage

	err = json.Unmarshal(d, &m)
	if err != nil {
		return nil, nil, err
	}

	other, err := mapOtherTransaction(bytes.NewBuffer(m))
	if err != nil {
		return nil, nil, err
	}

	t.Transaction.OtherTrans = other

	return &t.Meta, &t.Transaction, nil
}

type transferTransaction struct {
	base.CommonTransaction
	Amount    float64       `json:"amount,omitempty"`
	Recipient string        `json:"recipient,omitempty"`
	Message   base.Message  `json:"message,omitempty"`
	Signature string        `json:"signature,omitempty"`
	Mosaics   []base.Mosaic `json:"mosaics,omitempty"`
}

func (t *transferTransaction) toStruct() (base.Transaction, error) {
	return &base.TransferTransaction{
		CommonTransaction: base.CommonTransaction{
			Type:      t.Type,
			Version:   t.Version,
			Signer:    t.Signer,
			TimeStamp: t.TimeStamp,
			Fee:       t.Fee,
			Deadline:  t.Deadline,
		},
		Amount:    t.Amount,
		Recipient: t.Recipient,
		Message:   t.Message,
		Signature: t.Signature,
		Mosaics:   t.Mosaics,
	}, nil
}

// Each node can allow users to harvest with their delegated key on that node.
// The NIS configuration has entries for configuring the maximum number of allowed harvesters and optionally allow
// harvesting only for certain account addresses.
// The unlock info gives information about the maximum number of allowed harvesters and how many harvesters are
// already using the node.
type UnlockInfo struct {
	NumUnlocked int `json:"num-unlocked"`
	MaxUnlocked int `json:"max-unlocked"`
}

type Account struct {
	Account string `json:"account,omitempty"`
}

type HbAccountData struct {
	Accounts    *[]Account `json:"accounts,omitempty"`
	StartHeight int        `json:"startHeight,omitempty"`
	EndHeight   int        `json:"endHeight,omitempty"`
	IncrementBy int        `json:"incrementBy,omitempty"`
}

type HgAccountData struct {
	Address     string `json:"address,omitempty"`
	StartHeight int    `json:"startHeight,omitempty"`
	EndHeight   int    `json:"endHeight,omitempty"`
	IncrementBy int    `json:"incrementBy,omitempty"`
}

func MapTransactions(b *bytes.Buffer) ([]base.Transaction, error) {
	var wg sync.WaitGroup
	var err error

	var m []json.RawMessage

	err = json.Unmarshal(b.Bytes()[8:len(b.Bytes())-3], &m)
	if err != nil {
		return nil, err
	}

	txs := make([]base.Transaction, len(m))
	meta := make([]*TransactionMetaData, len(m))
	errs := make([]error, len(m))
	for i, t := range m {
		wg.Add(1)
		go func(i int, t json.RawMessage) {
			defer wg.Done()
			meta[i], txs[i], errs[i] = MapTransaction(bytes.NewBuffer(t))
		}(i, t)
	}
	wg.Wait()

	for _, err = range errs {
		if err != nil {
			return txs, err
		}
	}

	return txs, nil
}

func MapTransaction(b *bytes.Buffer) (*TransactionMetaData, base.Transaction, error) {
	var t uint16
	rawT := struct {
		Transaction struct {
			Type uint16
		}
	}{}

	err := json.Unmarshal(b.Bytes(), &rawT)
	if err != nil {
		return nil, nil, err
	}
	t = rawT.Transaction.Type

	if t == 0 {
		rawT := struct {
			Type uint16 `json:"type"`
		}{}

		err := json.Unmarshal(b.Bytes(), &rawT)
		if err != nil {
			return nil, nil, err
		}
		t = rawT.Type
	}

	var dto transactionDto = nil

	switch t {
	case model.Transfer:
		dto = &unconfirmedMosaicTransactionMetaDataPair{}
	case model.MultiSignTransaction:
		dto = &multiSignTransactionMetaDataPair{}
	default:
		fmt.Println(rawT.Transaction.Type)
	}

	return dtoToTransaction(b, dto)
}

func mapOtherTransaction(b *bytes.Buffer) (base.Transaction, error) {

	rawT := struct {
		Type uint16 `json:"type"`
	}{}

	err := json.Unmarshal(b.Bytes(), &rawT)
	if err != nil {
		return nil, err
	}

	var dto otherTransactionDto

	switch rawT.Type {
	case model.Transfer:
		dto = &transferTransaction{}
	default:
		fmt.Println(rawT.Type)
	}

	return dtoToOtherTransaction(b, dto)
}

type transactionDto interface {
	toStruct() (*TransactionMetaData, base.Transaction, error)
}

func dtoToTransaction(b *bytes.Buffer, dto transactionDto) (*TransactionMetaData, base.Transaction, error) {
	if dto == nil {
		return nil, nil, errors.New("dto can't be nil")
	}

	err := json.Unmarshal(b.Bytes(), dto)
	if err != nil {
		return nil, nil, err
	}

	meta, tx, err := dto.toStruct()
	if err != nil {
		return nil, nil, err
	}
	return meta, tx, nil
}

type otherTransactionDto interface {
	toStruct() (base.Transaction, error)
}

func dtoToOtherTransaction(b *bytes.Buffer, dto otherTransactionDto) (base.Transaction, error) {
	if dto == nil {
		return nil, errors.New("dto can't be nil")
	}

	err := json.Unmarshal(b.Bytes(), dto)
	if err != nil {
		return nil, err
	}

	tx, err := dto.toStruct()
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Gets the AccountMetaDataPair of an account.
// method Client - An Client endpoint struct point
// param {string} address - An account address
// return {struct} - An struct[AccountMetaDataPair] struct
// link http://bob.nem.ninja/docs/#accountMetaDataPair
func (c *Client) AccountData(address string) (AccountMetaDataPair, error) {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	c.URL.Path = "/account/get"
	req, err := c.buildReq(map[string]string{"address": address}, nil, http.MethodGet)
	if err != nil {
		return AccountMetaDataPair{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return AccountMetaDataPair{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return AccountMetaDataPair{}, err
	}

	var data AccountMetaDataPair
	if err := json.Unmarshal(byteArray, &data); err != nil {
		fmt.Println(err)
		return AccountMetaDataPair{}, err
	}
	return data, nil
}

// Gets the AccountMetaDataPair of an account with a public Key.
// method Client - An Client endpoint struct point
// param publicKey - An account public key
// return - An struct [AccountMetaDataPair] struct
// link http://bob.nem.ninja/docs/#accountMetaDataPair
func (c *Client) AccountDataFromPublicKey(publicKey string) (AccountMetaDataPair, error) {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	c.URL.Path = "/account/get/from-public-key"
	req, err := c.buildReq(map[string]string{"publicKey": publicKey}, nil, http.MethodGet)
	if err != nil {
		return AccountMetaDataPair{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return AccountMetaDataPair{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return AccountMetaDataPair{}, err
	}

	var data AccountMetaDataPair
	if err := json.Unmarshal(byteArray, &data); err != nil {
		fmt.Println(err)
		return AccountMetaDataPair{}, err
	}
	return data, nil
}

// Gets an array of harvest info objects for an account.
// method Client - An Client endpoint struct point
// param address - An account address
// return - An slice of [HarvestInfo] struct
// link http://bob.nem.ninja/docs/#harvestInfo
func (c *Client) HarvestedBlocks(address string) ([]HarvestInfo, error) {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	c.URL.Path = "/account/harvests"
	req, err := c.buildReq(map[string]string{"address": address}, nil, http.MethodGet)
	if err != nil {
		return []HarvestInfo{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return []HarvestInfo{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return []HarvestInfo{}, err
	}

	var data = struct{ Data []HarvestInfo }{}
	if err := json.Unmarshal(byteArray, &data); err != nil {
		return []HarvestInfo{}, err
	}
	return data.Data, nil
}

// Gets an array of TransactionMetaDataPair objects where the recipient has the address given as parameter to the request.
// method Client - An Client endpoint struct point
// param address - An account address
// param txHash - The 256 bit sha3 hash of the transaction up to which transactions are returned. (optional)
// param txId - The transaction id up to which transactions are returned. (optional)
// return - An slice of [TransactionMetaDataPair] struct
// link http://bob.nem.ninja/docs/#transactionMetaDataPair}
func (c *Client) IncomingTransactions(address, txHash, txId string) ([]TransactionMetaDataPair, error) {
	params := map[string]string{"address": address}
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	if txHash != "" {
		params["hash"] = txHash
	}
	if txId != "" {
		params["id"] = txId
	}

	c.URL.Path = "/account/transfers/incoming"
	req, err := c.buildReq(params, nil, http.MethodGet)
	if err != nil {
		return []TransactionMetaDataPair{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return []TransactionMetaDataPair{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return []TransactionMetaDataPair{}, err
	}

	var data = struct{ Data []TransactionMetaDataPair }{}
	if err := json.Unmarshal(byteArray, &data); err != nil {
		fmt.Println(err)
		return []TransactionMetaDataPair{}, err
	}
	return data.Data, nil
}

// Gets an array of TransactionMetaDataPair objects where the recipient has the address given as parameter to the request.
// method Client - An Client endpoint struct point
// param address - An account address
// param txHash - The 256 bit sha3 hash of the transaction up to which transactions are returned. (optional)
// param txId - The transaction id up to which transactions are returned. (optional)
// return - An slice of [TransactionMetaDataPair] struct
// link http://bob.nem.ninja/docs/#transactionMetaDataPair
func (c *Client) OutgoingTransactions(address, txHash, txId string) ([]TransactionMetaDataPair, error) {
	params := map[string]string{"address": address}
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	if txHash != "" {
		params["hash"] = txHash
	}
	if txId != "" {
		params["id"] = txId
	}

	c.URL.Path = "/account/transfers/outgoing"
	req, err := c.buildReq(params, nil, http.MethodGet)
	if err != nil {
		return []TransactionMetaDataPair{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return []TransactionMetaDataPair{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return []TransactionMetaDataPair{}, err
	}

	var data = struct{ Data []TransactionMetaDataPair }{}
	if err := json.Unmarshal(byteArray, &data); err != nil {
		fmt.Println(err)
		return []TransactionMetaDataPair{}, err
	}
	return data.Data, nil
}

// Gets the array of transactions for which an account is the sender or receiver and which
// have not yet been included in a block.
// method Client - An Client endpoint struct point
// param address - An account address
// return - An slice of [UnconfirmedTransactionMetaDataPair] struct
// link http://bob.nem.ninja/docs/#unconfirmedTransactionMetaDataPair
func (c *Client) UnconfirmedTransactions(address string) ([]base.Transaction, error) {

	b := new(bytes.Buffer)
	params := map[string]string{"address": address}
	timeout := 10 * time.Second
	client := http.Client{
		Timeout: timeout,
	}
	c.URL.Path = "/account/unconfirmedTransactions"
	req, err := c.buildReq(params, nil, http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		b := &bytes.Buffer{}
		_, _ = b.ReadFrom(resp.Body)
		return nil, errors.New(b.String())
	}

	_, _ = io.Copy(b, resp.Body)

	return MapTransactions(b)
}

// Gets information about the maximum number of allowed harvesters and
// how many harvesters are already using the node
// method Client - An Client endpoint struct point
// return - An [UnlockInfo] struct
// link http://bob.nem.ninja/docs/#retrieving-the-unlock-info
func (c *Client) UnlockInfo() (UnlockInfo, error) {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	c.URL.Path = "/account/unlocked/info"
	req, err := c.buildReq(nil, nil, http.MethodPost)
	if err != nil {
		return UnlockInfo{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return UnlockInfo{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	var data = UnlockInfo{}
	if err := json.Unmarshal(byteArray, &data); err != nil {
		fmt.Println(err)
		return UnlockInfo{}, err
	}
	return data, nil
}

// Unlocks an account (starts harvesting).
// method Client - An Client endpoint struct point
// param privateKey - A delegated account private key
// return - error
func (c *Client) StartHarvesting(privateKey string) error {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	payload, err := json.Marshal(map[string]string{"value": privateKey})
	if err != nil {
		return err
	}

	c.URL.Path = "/account/unlock"
	req, err := c.buildReq(nil, payload, http.MethodPost)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return err
	}
	return nil
}

// Locks an account (stops harvesting).
// method Client - An Client endpoint struct point
// param privateKey - A delegated account private key
// return - error
func (c *Client) StopHarvesting(privateKey string) error {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	payload, err := json.Marshal(map[string]string{"value": privateKey})
	if err != nil {
		return err
	}

	c.URL.Path = "/account/lock"
	req, err := c.buildReq(nil, payload, http.MethodPost)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return err
	}
	return nil
}

// Gets the AccountMetaDataPair of the account for which the given account is the delegate account
// method Client - An Client endpoint struct point
// param address - An account address
// return - An struct[AccountMetaDataPair] struct
// link http://bob.nem.ninja/docs/#accountMetaDataPair
func (c *Client) Forwarded(address string) (AccountMetaDataPair, error) {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	c.URL.Path = "/account/get/forwarded"
	req, err := c.buildReq(map[string]string{"address": address}, nil, http.MethodGet)
	if err != nil {
		return AccountMetaDataPair{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return AccountMetaDataPair{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return AccountMetaDataPair{}, err
	}

	var data AccountMetaDataPair
	if err := json.Unmarshal(byteArray, &data); err != nil {
		fmt.Println(err)
		return AccountMetaDataPair{}, err
	}
	return data, nil
}

// Gets namespaces that an account owns
// method Client - An Client endpoint struct point
// param address - An account address
// param parent - The namespace parent (optional)
// return - An slice of [Namespace] struct
// link http://bob.nem.ninja/docs/#namespaceMetaDataPair
func (c *Client) NamespacesOwned(address, parent string) ([]Namespace, error) {
	params := map[string]string{"address": address}
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	if parent != "" {
		params["parent"] = parent
	}
	c.URL.Path = "/account/namespace/page"
	req, err := c.buildReq(params, nil, http.MethodGet)
	if err != nil {
		return []Namespace{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return []Namespace{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return []Namespace{}, err
	}

	var data = struct {
		Data []Namespace
	}{}
	if err := json.Unmarshal(byteArray, &data); err != nil {
		fmt.Println(err)
		return []Namespace{}, err
	}
	if data.Data == nil {
		return []Namespace{}, err
	}

	return data.Data, nil
}

// Gets mosaic definitions that an account has created
// method Client - An Client endpoint struct point
// param address - An account address
// param parent - The namespace parent (optional)
// return - An slice of [MosaicDefinition] struct
// link http://bob.nem.ninja/docs/#mosaicDefinition
func (c *Client) MosaicDefinitionsCreated(address, parent string) ([]base.MosaicDefinition, error) {
	params := map[string]string{"address": address}
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	if parent != "" {
		params["parent"] = parent
	}
	c.URL.Path = "/account/mosaic/definition/page"
	req, err := c.buildReq(params, nil, http.MethodGet)
	if err != nil {
		return []base.MosaicDefinition{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return []base.MosaicDefinition{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return []base.MosaicDefinition{}, err
	}

	var data = struct {
		Data []base.MosaicDefinition
	}{}
	if err := json.Unmarshal(byteArray, &data); err != nil {
		fmt.Println(err)
		return []base.MosaicDefinition{}, err
	}
	if data.Data == nil {
		return []base.MosaicDefinition{}, err
	}

	return data.Data, nil
}

// Gets mosaic definitions that an account owns
// method Client - An Client endpoint struct point
// param address - An account address
// return - An slice of [MosaicDefinition] struct
// link http://bob.nem.ninja/docs/#mosaicDefinition
func (c *Client) MosaicDefinitionsOwned(address string) ([]base.MosaicDefinition, error) {
	params := map[string]string{"address": address}
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	c.URL.Path = "/account/mosaic/owned/definition"
	req, err := c.buildReq(params, nil, http.MethodGet)
	if err != nil {
		return []base.MosaicDefinition{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return []base.MosaicDefinition{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return []base.MosaicDefinition{}, err
	}

	var data = struct {
		Data []base.MosaicDefinition
	}{}
	if err := json.Unmarshal(byteArray, &data); err != nil {
		fmt.Println(err)
		return []base.MosaicDefinition{}, err
	}
	if data.Data == nil {
		return []base.MosaicDefinition{}, err
	}

	return data.Data, nil
}

// Gets mosaics that an account owns
// method Client - An Client endpoint struct point
// param address - An account address
// return - An slice of [Mosaic] struct
// link http://bob.nem.ninja/docs/#mosaic
func (c *Client) MosaicsOwned(address string) ([]base.Mosaic, error) {
	params := map[string]string{"address": address}
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	c.URL.Path = "/account/mosaic/owned"
	req, err := c.buildReq(params, nil, http.MethodGet)
	if err != nil {
		return []base.Mosaic{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return []base.Mosaic{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return []base.Mosaic{}, err
	}

	var data = struct {
		Data []base.Mosaic
	}{}
	if err := json.Unmarshal(byteArray, &data); err != nil {
		fmt.Println(err)
		return []base.Mosaic{}, err
	}
	if data.Data == nil {
		return []base.Mosaic{}, err
	}

	return data.Data, nil
}

// Gets all transactions of an account
// method Client - An Client endpoint struct point
// param address - An account address
// param txHash - The 256 bit sha3 hash of the transaction up to which transactions are returned. (optional)
// param txId - The transaction id up to which transactions are returned. (optional)
// return - An slice of [TransactionMetaDataPair] struct
// link http://bob.nem.ninja/docs/#transactionMetaDataPair
func (c *Client) AllTransactions(address, txHash, txId string) ([]TransactionMetaDataPair, error) {
	params := map[string]string{"address": address}
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	if txHash != "" {
		params["hash"] = txHash
	}
	if txId != "" {
		params["id"] = txId
	}

	c.URL.Path = "/account/transfers/all"
	req, err := c.buildReq(params, nil, http.MethodGet)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return []TransactionMetaDataPair{}, err
	}

	var data = struct{ Data []TransactionMetaDataPair }{}
	if err := json.Unmarshal(byteArray, &data); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data.Data, nil
}

// Gets the AccountMetaDataPair of an array of accounts.
// method Client - An Client endpoint struct point
// param addresses - An array of account addresses
// return - An slice that contains an array of [AccountMetaDataPair] struct
// link http://bob.nem.ninja/docs/#accountMetaDataPair
func (c Client) GetBatchAccountData(addresses []string) ([]AccountMetaDataPair, error) {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	var payloadBuilder []map[string]string
	for _, address := range addresses {
		payloadBuilder = append(payloadBuilder, map[string]string{"account": address})
	}
	payload, err := json.Marshal(map[string][]map[string]string{"data": payloadBuilder})
	if err != nil {
		return []AccountMetaDataPair{}, err
	}
	c.URL.Path = "/account/get/batch"
	req, err := c.buildReq(nil, payload, http.MethodPost)
	if err != nil {
		return []AccountMetaDataPair{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return []AccountMetaDataPair{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return []AccountMetaDataPair{}, err
	}
	// The data is returned as a nested json array
	// This enables us to not return the array nested
	// as a value under a "data" key
	data := struct{ Data []AccountMetaDataPair }{}
	if err = json.Unmarshal(byteArray, &data); err != nil {
		return []AccountMetaDataPair{}, err
	}
	return data.Data, nil
}

// Gets the AccountMetaDataPair of an array of accounts from an historical height.
// method Client - An Client endpoint struct point
// param addresses - An array of account addresses
// param block - The block height
// return - An slice Account information for all the accounts on the given block
func (c Client) GetBatchHistoricalAccountData(addresses []string, block int) ([]AccountMetaDataPair, error) {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	var Accounts []Account

	for _, address := range addresses {
		Accounts = append(Accounts, Account{Account: address})
	}
	Ojt := HbAccountData{}
	Ojt.Accounts = &Accounts
	Ojt.StartHeight = block
	Ojt.EndHeight = block
	Ojt.IncrementBy = 1

	//Ojt = append(Ojt, map[string]string{"accounts": Accounts})

	payload, err := json.Marshal(Ojt)
	if err != nil {
		return []AccountMetaDataPair{}, err
	}

	c.URL.Path = "/account/historical/get/batch"
	req, err := c.buildReq(nil, payload, http.MethodPost)
	if err != nil {
		return []AccountMetaDataPair{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return []AccountMetaDataPair{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return []AccountMetaDataPair{}, err
	}
	// The data is returned as a nested json array
	// This enables us to not return the array nested
	// as a value under a "data" key
	data := struct{ Data []AccountMetaDataPair }{}
	if err = json.Unmarshal(byteArray, &data); err != nil {
		return []AccountMetaDataPair{}, err
	}
	return data.Data, nil
}

// Gets the AccountMetaDataPair of an account from a certain block.
// method Client - An Client endpoint struct point
// param address - An account address
// param block - the block height
// return - An slice [AccountMetaDataPair] struct
// link http://bob.nem.ninja/docs/#accountMetaDataPair
func (c Client) GetHistoricalAccountData(addresses string, block int) ([]AccountMetaDataPair, error) {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	params := map[string]string{"address": addresses}

	bck := fmt.Sprintf("%v", block)
	params["startHeight"] = bck
	params["endHeight"] = bck
	params["incrementBy"] = "1"

	c.URL.Path = "/account/historical/get"
	req, err := c.buildReq(params, nil, http.MethodGet)
	if err != nil {
		return []AccountMetaDataPair{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return []AccountMetaDataPair{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return []AccountMetaDataPair{}, err
	}
	// The data is returned as a nested json array
	// This enables us to not return the array nested
	// as a value under a "data" key
	data := struct{ Data []AccountMetaDataPair }{}
	if err = json.Unmarshal(byteArray, &data); err != nil {
		return []AccountMetaDataPair{}, err
	}
	return data.Data, nil
}
