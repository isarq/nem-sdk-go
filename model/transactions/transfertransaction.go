package transactions

import (
	"errors"
	"github.com/isarq/nem-sdk-go/base"
	"github.com/isarq/nem-sdk-go/extras"
	"github.com/isarq/nem-sdk-go/model"
	"github.com/isarq/nem-sdk-go/utils"

	"math"
	"strings"
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

func (t *Transfer) GetTx() base.Transaction {
	return base.Transaction{}
}

// Prepare a transfer transaction struct
// param common - A common struct
// param r - An un-prepared TransferTransaction method
// param network - A network id
// return - A [TransferTransaction] struct
// link http://bob.nem.ninja/docs/#transferTransaction
func (r *Transfer) Prepare(common Common, network int) base.TxDict {
	var msc txPrepare
	if extras.IsEmpty(common) || extras.IsEmpty(network) {
		err := errors.New("missing parameter !")
		panic(err)
	}
	kp := model.KeyPairCreate(common.PrivateKey)
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

	if network == model.Data.Testnet.ID {
		msc.due = 60
	} else {
		msc.due = 24 * 60
	}
	msc.network = network

	rt := constructtx(msc)
	if r.IsMultisig && r.MultisigAccount != "" {
		return MultisigWrapper(kp.PublicString(), rt, msc.due, network)
	}
	return rt
}

// Prepare a mosaic transfer transaction struct
// param common - A common struct
// param tx - The un-prepared transfer transaction struct
// param mosaicDefinitionMetaDataPair - The mosaicDefinitionMetaDataPair object with properties of mosaics to send
// param network - A network id
// return - A [TransferTransaction] struct ready for serialization
// link http://bob.nem.ninja/docs/#transferTransaction
func (r *Transfer) PrepareMosaic(common Common, network int) base.TxDict {
	var msc txPrepare
	if extras.IsEmpty(common) || extras.IsEmpty(network) {
		err := errors.New("missing parameter !")
		panic(err)
	}
	kp := model.KeyPairCreate(common.PrivateKey)
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

	if network == model.Data.Testnet.ID {
		msc.due = 60
	} else {
		msc.due = 24 * 60
	}
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
		TimeStamp: data.TimeStamp,
		Amount:    msc.amount,
		Fee:       totalFee,
		Recipient: strings.ToUpper(strings.Replace(msc.recipientCompressedKey, "-", "", -1)),
		Type:      data.Type,
		Deadline:  data.Deadline,
		Message:   msc.message,
		Version:   data.Version,
		Signer:    data.Signer,
		Mosaics:   msc.mosaics,
	}
	return &custom
}
