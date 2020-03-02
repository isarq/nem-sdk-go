package transactions

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/isarq/nem-sdk-go/base"
	"github.com/isarq/nem-sdk-go/com/requests"
	"github.com/isarq/nem-sdk-go/extras"
	"github.com/isarq/nem-sdk-go/model"
	"github.com/isarq/nem-sdk-go/utils"
)

type txPrepare struct {
	senderPublicKey        string
	recipientCompressedKey string
	amount                 float64
	message                base.Message
	msgFee                 float64
	due                    int64
	mosaics                []base.Mosaic
	mosaicsFee             float64
	network                int
}

type Transfer struct {
	Amount             float64       `json:"amount"`
	Recipient          string        `json:"recipient"`
	RecipientPublicKey string        `json:"recipientPublicKey"`
	IsMultisig         bool          `json:"isMultisig"`
	MultisigAccount    string        `json:"multisigAccount"`
	Message            string        `json:"message"`
	MessageType        int           `json:"messageType"`
	Mosaics            []base.Mosaic `json:"mosaics"`
}

func (r *Transfer) Get() {
	r.MultisigAccount = ""
	r.IsMultisig = false
}

func (t *Transfer) GetType() int {
	return 0
}

// Prepare a transfer transaction struct
// param common - A common struct
// param r - An un-prepared TransferTransaction method
// param network - A network id
// return - A [TransferTransaction] struct
// link http://bob.nem.ninja/docs/#transferTransaction
func (r *Transfer) Prepare(common Common, network int) (base.Transaction, error) {
	var msc txPrepare
	if extras.IsEmpty(common) || extras.IsEmpty(network) {
		return nil, errors.New("missing parameter !")
	}
	kp, err := model.KeyPairCreate(common.PrivateKey)
	if err != nil {
		return nil, err
	}
	if r.IsMultisig {
		if r.MultisigAccount != "" {
			if !utils.IsPublicKeyValid(r.MultisigAccount) {
				return nil, errors.New("Invalid public key!")
			}
			msc.senderPublicKey = r.MultisigAccount
		} else {
			return nil, errors.New("must place a publickey of the multifirm account")
		}
	} else {
		msc.senderPublicKey = kp.PublicString()
	}

	msc.recipientCompressedKey = strings.ToUpper(strings.Replace(r.Recipient, "-", "", -1))

	msc.amount = math.Round(r.Amount * 1000000)

	msc.message = MsgPrepare(common, r)

	msc.msgFee = model.CalculateMessage(msc.message, false)

	if network == model.Data.Testnet.ID {
		msc.due = 60
	} else {
		msc.due = 24 * 60
	}
	msc.network = network

	rt := constructtx(msc)
	if r.IsMultisig && r.MultisigAccount != "" {
		return MultisigWrapper(kp.PublicString(), rt, msc.due, network), nil
	}
	return rt, nil
}

// Prepare a mosaic transfer transaction struct
// param common - A common struct
// param tx - The un-prepared transfer transaction struct
// param mosaicDefinitionMetaDataPair - The mosaicDefinitionMetaDataPair object with properties of mosaics to send
// param network - A network id
// return - A [TransferTransaction] struct ready for serialization
// link http://bob.nem.ninja/docs/#transferTransaction
func (r *Transfer) PrepareMosaic(common Common, mosaicDefinitionMetaDataPair map[string]base.MosaicDefinition,
	client *requests.Client, network int) base.Transaction {
	supplys := make(map[string]float64)
	var msc txPrepare
	if extras.IsEmpty(common) || extras.IsEmpty(network) || extras.IsEmpty(mosaicDefinitionMetaDataPair) {
		err := errors.New("missing parameter !")
		panic(err)
	}
	kp, _ := model.KeyPairCreate(common.PrivateKey)
	if r.IsMultisig {
		if r.MultisigAccount != "" {
			if !utils.IsPublicKeyValid(r.MultisigAccount) {
				panic(nil)
			}
			msc.senderPublicKey = r.MultisigAccount
		} else {
			err := errors.New("must place a publickey of the multifirm account")
			panic(err)
		}
	} else {
		msc.senderPublicKey = kp.PublicString()
	}

	msc.recipientCompressedKey = strings.ToUpper(strings.Replace(r.Recipient, "-", "", -1))

	msc.amount = math.Round(r.Amount * 1000000)

	msc.message = MsgPrepare(common, r)

	msc.msgFee = model.CalculateMessage(msc.message, false)

	//Gets the current supply of a mosaic
	for _, b := range r.Mosaics {
		fullMosaicName := utils.MosaicIdToName(b.MosaicID)
		supply, err := client.Supply(fullMosaicName)
		if err != nil {
			fmt.Println(utils.Struc2Json(err))
		}
		supplys[fullMosaicName] = float64(supply.Supply)
	}

	msc.mosaicsFee = model.CalculateMosaics(msc.amount, mosaicDefinitionMetaDataPair, r.Mosaics, supplys)

	if network == model.Data.Testnet.ID {
		msc.due = 60
	} else {
		msc.due = 24 * 60
	}
	msc.mosaics = r.Mosaics

	//msc.mosaicsFee = model.Ca

	msc.network = network

	rt := constructtx(msc)
	if r.IsMultisig && r.MultisigAccount != "" {
		return MultisigWrapper(kp.PublicString(), rt, msc.due, network)
	}
	return rt
}

// Create a namespace provision transaction struct
// param msc network - A nsPrepare struct
// return - A [ProvisionNamespaceTransaction] struct
// link http://bob.nem.ninja/docs/#provisionNamespaceTransaction
func constructtx(msc txPrepare) *base.TransferTransaction {
	timeStamp := utils.CreateNEMTimeStamp()
	var version int
	if extras.IsEmpty(msc.mosaics) {
		version = model.GetVersion(1, msc.network)
	} else {
		version = model.GetVersion(2, msc.network)
	}
	data := CommonPart(model.Transfer, version, timeStamp, msc.due, msc.senderPublicKey)
	var fee float64
	if !extras.IsEmpty(msc.mosaics) {
		fee = msc.mosaicsFee
	} else {
		fee = model.CurrentFeeFactor * model.CalculateMinimum(msc.amount/1000000)
	}
	totalFee := math.Floor((msc.msgFee + fee) * 1000000)

	custom := base.TransferTransaction{
		CommonTransaction: base.CommonTransaction{
			TimeStamp: data.TimeStamp,
			Fee:       totalFee,
			Version:   data.Version,
			Signer:    data.Signer,
			Type:      data.Type,
			Deadline:  data.Deadline,
		},
		Amount:    msc.amount,
		Recipient: strings.ToUpper(strings.Replace(msc.recipientCompressedKey, "-", "", -1)),

		Message: msc.message,

		Mosaics: msc.mosaics,
	}
	return &custom
}
