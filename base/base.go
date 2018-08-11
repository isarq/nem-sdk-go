package base

type Error struct {
	TimeStamp int
	Error     string
	Message   string
	Status    int
}

// Before a mosaic can be created or transferred, a corresponding
// definition of the mosaic has to be created and published to the network.
// This is done via a mosaic definition creation transaction.
type MosaicDefinitionCreationTransaction struct {
	TimeStamp        int64            `json:"timeStamp"`
	Fee              float64          `json:"fee"`
	Type             int              `json:"type"`
	Deadline         int64            `json:"deadline"`
	Version          int              `json:"version"`
	Signer           string           `json:"signer"`
	CreationFee      float64          `json:"creationFee"`
	CreationFeeSink  string           `json:"creationFeeSink"`
	MosaicDefinition MosaicDefinition `json:"mosaicDefinition"`
}

// A mosaic definition describes an asset class. Some fields are mandatory
// while others are optional. The properties of a mosaic definition always
// have a default value and only need to be supplied if they differ from the default value.
type MosaicDefinition struct {
	Creator     string       `json:"creator"`
	Description string       `json:"description"`
	ID          MosaicID     `json:"id,omitempty"`
	Properties  []Properties `json:"properties"`
	Levy        Levy         `json:"levy,omitempty"`
}

type MosaicID struct {
	NamespaceID string `json:"namespaceId,omitempty"`
	Name        string `json:"name,omitempty"`
}

// A mosaic describes an instance of a mosaic definition.
// Mosaics can be transferred by means of a transfer transaction.
type Mosaic struct {
	MosaicID MosaicID `json:"mosaicID,omitempty"`
	Quantity float64  `json:"quantity,omitempty"`
}

type Properties struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Levy struct {
	Type      int      `json:"type,omitempty"`
	Recipient string   `json:"recipient,omitempty"`
	MosaicID  MosaicID `json:"mosaicId,omitempty"`
	Fee       float64  `json:"fee,omitempty"`
}

type Node struct {
	Host string
	Port int
}

type MessageType struct {
	Value int
	Name  string
}

type ConsModif struct {
	ModificationType   int
	CosignatoryAccount string
}

type TransferTransaction struct {
	TimeStamp int64    `json:"timeStamp,omitempty"`
	Amount    float64  `json:"amount,omitempty"`
	Signature string   `json:"signature,omitempty"`
	Fee       float64  `json:"fee,omitempty"`
	Recipient string   `json:"recipient,omitempty"`
	Type      int      `json:"type,omitempty"`
	Deadline  int64    `json:"deadline,omitempty"`
	Message   Message  `json:"message,omitempty"`
	Version   int      `json:"version,omitempty"`
	Signer    string   `json:"signer,omitempty"`
	Mosaics   []Mosaic `json:"mosaics,omitempty"`
}

type Common struct {
	Password, PrivateKey string
	IsHW                 bool
}

// Chain is the type of NEM chain.
type Chain struct {
	ID, Prefix int
	Char       string
}

type Data struct {
	Testnet, Mainnet, Mijin Chain
}

type Message struct {
	Type      int    `json:"type,omitempty"`
	Payload   string `json:"payload,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
}

type SignatureT struct {
	OtherHash struct {
		Data string
	}
	OtherAccount string
}

type Supply struct {
	Mosaic          string `json:"mosaic"`
	SupplyType      int    `json:"supplyType"`
	Delta           int    `json:"delta"`
	IsMultisig      bool   `json:"isMultisig"`
	MultisigAccount string `json:"multisigAccount"`
}

type MultisigAggregateModific struct {
	Modifications   []interface{} `json:"modifications"`
	RelativeChange  interface{}   `json:"relativeChange"`
	IsMultisig      bool          `json:"isMultisig"`
	MultisigAccount string        `json:"multisigAccount"`
}

type ImportanceTransfer struct {
	RemoteAccount   string `json:"remoteAccount"`
	Mode            int    `json:"mode"`
	IsMultisig      bool   `json:"isMultisig"`
	MultisigAccount string `json:"multisigAccount"`
}

type CommonTransaction struct {
	Type      int
	Version   int
	Signer    string
	TimeStamp int64
	Fee       float64
	Deadline  int64
}

type ProvisionNamespaceTransaction struct {
	TimeStamp     int64   `json:"timeStamp"`
	Fee           float64 `json:"fee"`
	Type          int     `json:"type"`
	Deadline      int64   `json:"deadline"`
	Version       int     `json:"version"`
	Signer        string  `json:"signer"`
	RentalFeeSink string  `json:"rentalFeeSink"`
	RentalFee     float64 `json:"rentalFee"`
	NewPart       string  `json:"newPart"`
	Parent        string  `json:"parent"`
}

type MultisigSignatureTransaction struct {
	TimeStamp int    `json:"timeStamp"`
	Fee       int    `json:"fee"`
	Type      int    `json:"type"`
	Deadline  int    `json:"deadline"`
	Version   int    `json:"version"`
	Signer    string `json:"signer"`
	OtherHash struct {
		Data string `json:"data"`
	} `json:"otherHash"`
	OtherAccount string `json:"otherAccount"`
}

type TransactionResponce struct {
	TimeStamp  int64                          `json:"timeStamp"`
	Amount     float64                        `json:"amount"`
	Fee        float64                        `json:"fee"`
	Recipient  string                         `json:"recipient,omitempty"`
	Type       int                            `json:"type,omitempty"`
	Deadline   int64                          `json:"deadline"`
	Message    *Message                       `json:"message,omitempty"`
	Version    int                            `json:"version,omitempty"`
	Signer     string                         `json:"signer,omitempty"`
	OtherTrans Transaction                    `json:"otherTrans,omitempty"`
	Signatures []MultisigSignatureTransaction `json:"signatures,omitempty"`
}

type MultisigTransaction struct {
	TimeStamp  int64                          `json:"timeStamp"`
	Fee        float64                        `json:"fee"`
	Type       int                            `json:"type"`
	Deadline   int64                          `json:"deadline"`
	Version    int                            `json:"version"`
	Signer     string                         `json:"signer"`
	OtherTrans interface{}                    `json:"otherTrans"`
	Signatures []MultisigSignatureTransaction `json:"signatures,omitempty"`
}

type Transaction struct {
	TimeStamp int64    `json:"timeStamp,omitempty"`
	Amount    float64  `json:"amount,omitempty"`
	Fee       float64  `json:"fee,omitempty"`
	Recipient string   `json:"recipient,omitempty"`
	Type      int      `json:"type,omitempty"`
	Deadline  int64    `json:"deadline,omitempty"`
	Message   *Message `json:"message,omitempty"`
	Version   int      `json:"version,omitempty"`
	Signer    string   `json:"signer,omitempty"`
}

type Tx interface {
	GetType() int
	GetTx() Transaction
}

type TxDict interface {
	GetCommon() CommonTransaction
	//GetMosaic() MosaicDefinition
	//GetMosaicId() MosaicID
	//GetMosaicTx() *MosaicDefinitionCreationTransaction
}

func (t *MosaicDefinitionCreationTransaction) GetType() int {
	mosaic := t.Type
	return mosaic
}

func (t *MosaicDefinitionCreationTransaction) GetCommon() CommonTransaction {
	return CommonTransaction{
		Type:      t.Type,
		Version:   t.Version,
		TimeStamp: t.TimeStamp,
		Deadline:  t.Deadline,
		Signer:    t.Signer,
		Fee:       t.Fee,
	}
}

func (t *MosaicDefinitionCreationTransaction) GetMosaic() MosaicDefinition {
	mosaic := t.MosaicDefinition
	return mosaic
}

func (t *MosaicDefinitionCreationTransaction) GetMosaicId() MosaicID {
	mosaic := t.MosaicDefinition.ID
	return mosaic
}

func (t *MosaicDefinitionCreationTransaction) GetMosaicTx() *MosaicDefinitionCreationTransaction {
	return t
}

func (t *MosaicDefinitionCreationTransaction) GetTx() Transaction {
	return Transaction{
		TimeStamp: t.TimeStamp,
		Amount:    0,
		Fee:       t.Fee,
		Recipient: "",
		Type:      t.Type,
		Deadline:  t.Deadline,
		Version:   t.Version,
		Signer:    t.Signer,
	}
}

func (t *ProvisionNamespaceTransaction) GetType() int {
	return t.Type
}

func (t *ProvisionNamespaceTransaction) GetCommon() CommonTransaction {
	return CommonTransaction{
		Type:      t.Type,
		Version:   t.Version,
		TimeStamp: t.TimeStamp,
		Deadline:  t.Deadline,
		Signer:    t.Signer,
		Fee:       t.Fee,
	}
}

func (t *ProvisionNamespaceTransaction) GetTx() Transaction {
	return Transaction{
		TimeStamp: t.TimeStamp,
		Amount:    0,
		Fee:       t.Fee,
		Recipient: "",
		Type:      t.Type,
		Deadline:  t.Deadline,
		Version:   t.Version,
		Signer:    t.Signer,
	}
}

func (t *MultisigTransaction) GetType() int {
	return t.Type
}

func (t *MultisigTransaction) GetCommon() CommonTransaction {
	return CommonTransaction{
		Type:      t.Type,
		Version:   t.Version,
		TimeStamp: t.TimeStamp,
		Deadline:  t.Deadline,
		Signer:    t.Signer,
		Fee:       t.Fee,
	}
}

func (t *MultisigTransaction) GetTx() Transaction {
	tx := t.OtherTrans.(Transaction)
	return Transaction{
		TimeStamp: tx.TimeStamp,
		Amount:    tx.Amount,
		Fee:       tx.Fee,
		Recipient: tx.Recipient,
		Type:      tx.Type,
		Deadline:  tx.Deadline,
		Version:   tx.Version,
		Signer:    tx.Signer,
		Message:   tx.Message,
	}
}

func (t *TransferTransaction) GetType() int {
	return t.Type
}

func (t *TransferTransaction) GetCommon() CommonTransaction {
	return CommonTransaction{
		Type:      t.Type,
		Version:   t.Version,
		TimeStamp: t.TimeStamp,
		Deadline:  t.Deadline,
		Signer:    t.Signer,
		Fee:       t.Fee,
	}
}

func (t *TransferTransaction) GetTx() Transaction {
	return Transaction{
		TimeStamp: t.TimeStamp,
		Amount:    0,
		Fee:       t.Fee,
		Recipient: t.Recipient,
		Type:      t.Type,
		Deadline:  t.Deadline,
		Version:   t.Version,
		Signer:    t.Signer,
	}
}

func (t *Transaction) GetType() int {
	return t.Type
}

func (t *Transaction) GetCommon() CommonTransaction {
	return CommonTransaction{
		Type:      t.Type,
		Version:   t.Version,
		TimeStamp: t.TimeStamp,
		Deadline:  t.Deadline,
		Signer:    t.Signer,
		Fee:       t.Fee,
	}
}

func (t *Transaction) GetTx() Transaction {
	return Transaction{
		TimeStamp: t.TimeStamp,
		Amount:    0,
		Fee:       t.Fee,
		Recipient: t.Recipient,
		Type:      t.Type,
		Deadline:  t.Deadline,
		Version:   t.Version,
		Signer:    t.Signer,
	}
}
