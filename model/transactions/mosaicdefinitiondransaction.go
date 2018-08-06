package transactions

import (
	"github.com/pkg/errors"
	"github.com/isarq/nem-sdk-go/base"
	"github.com/isarq/nem-sdk-go/extras"
	"github.com/isarq/nem-sdk-go/model"
	"github.com/isarq/nem-sdk-go/utils"
	"strings"
)

type TxDict interface {
	Get() interface{}
	Create() interface{}
}

// A mosaic definition describes an asset class. Some fields are mandatory
// while others are optional. The properties of a mosaic definition always
// have a default value and only need to be supplied if they differ from the default value.
type MosaicDefinition struct {
	MosaicName      string `json:"mosaicName,omitempty"`
	NamespaceParent struct {
		Fqn string `json:"fqn,omitempty"`
	} `json:"namespaceParent,omitempty"`
	MosaicDescription string            `json:"mosaicDescription,omitempty"`
	Properties        []base.Properties `json:"properties,omitempty"`
	Levy              Levy              `json:"levy,omitempty"`
	IsMultisig        bool              `json:"isMultisig"`
	MultisigAccount   string            `json:"multisigAccount"`
}

type Levy struct {
	Mosaic struct {
		NamespaceID string `json:"namespaceId,omitempty"`
		Name        string `json:"name,omitempty"`
	} `json:"mosaic,omitempty"`
	Address string  `json:"address,omitempty"`
	FeeType int     `json:"feeType,omitempty"`
	Fee     float64 `json:"fee,omitempty"`
}

type mosaicPrepare struct {
	senderPublicKey   string
	rentalFeeSink     string
	rentalFee         float64
	namespaceParent   string
	mosaicName        string
	mosaicDescription string
	mosaicProperties  []base.Properties
	levy              Levy
	due               int64
	network           int
}

func (r *MosaicDefinition) Get() {
	r.Properties = []base.Properties{
		{Name: "supplyMutable", Value: "true"},
		{Name: "transferable", Value: "true"},
	}
	r.Levy.FeeType = 1
	r.Levy.Fee = 5
	r.IsMultisig = false
}

// Prepare a mosaic definition transaction
// argument	r - An un-prepared mosaicDefinitionTransaction struct
// param common - A common struct
// param network - A network id
// return A [MosaicDefinitionCreationTransaction] struc ready for serialization
// link http://bob.nem.ninja/docs/#mosaicDefinitionCreationTransaction
func (r MosaicDefinition) Prepare(common Common, network int) *base.MosaicDefinitionCreationTransaction {
	var msc mosaicPrepare
	if !utils.IsPrivateKeyValid(common.PrivateKey) {
		panic(nil)
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
	msc.rentalFeeSink = strings.ToUpper(strings.Replace(model.Mosaic[network], "-", "", -1))
	//fmt.Println("CreationFeeSink: ", msc.rentalFeeSink)
	msc.rentalFee = model.MosaicDefinitionTransaction
	//fmt.Printf("RentalFee: %v", msc.rentalFee)
	msc.namespaceParent = r.NamespaceParent.Fqn
	msc.mosaicName = r.MosaicName
	msc.mosaicDescription = r.MosaicDescription
	msc.mosaicProperties = r.Properties
	if !extras.IsEmpty(r.Levy.Mosaic) {
		msc.levy = r.Levy
	}

	if network == model.Data.Testnet.ID {
		msc.due = 60
	} else {
		msc.due = 24 * 60
	}
	msc.network = network
	return constructMs(msc)
}

// Create a mosaic definition transaction struct
// param msc - A mosaicPrepare struct
// return - A [MosaicDefinitionCreationTransaction] struct
// link http://bob.nem.ninja/docs/#mosaicDefinitionCreationTransaction
func constructMs(msc mosaicPrepare) *base.MosaicDefinitionCreationTransaction {
	timeStamp := utils.CreateNEMTimeStamp()
	version := model.GetVersion(1, msc.network)
	data := CommonPart(model.Mosaicdefinition, version, timeStamp, msc.due, msc.senderPublicKey)
	fee := model.NamespaceAndMosaicCommon
	levyData := base.Levy{}
	if !extras.IsEmpty(msc.levy) {
		levyData.Type = msc.levy.FeeType
		levyData.Recipient = msc.levy.Address
		levyData.MosaicID = msc.levy.Mosaic
		levyData.Fee = msc.levy.Fee
	}
	custom := base.MosaicDefinitionCreationTransaction{
		TimeStamp:       data.TimeStamp,
		Signer:          data.Signer,
		Type:            data.Type,
		Deadline:        data.Deadline,
		Version:         data.Version,
		CreationFeeSink: msc.rentalFeeSink,
		CreationFee:     msc.rentalFee,
		MosaicDefinition: base.MosaicDefinition{
			Creator: msc.senderPublicKey,
			ID: struct {
				NamespaceID string `json:"namespaceId,omitempty"`
				Name        string `json:"name,omitempty"`
			}{
				NamespaceID: msc.namespaceParent,
				Name:        msc.mosaicName},
			Description: msc.mosaicDescription,
			Properties:  msc.mosaicProperties,
			Levy:        levyData,
		},
		Fee: fee,
	}
	return &custom
}

// The common part of transactions
// param txType - A type of transaction
// param senderPublicKey - A sender public key
// param timeStamp - A timestamp for the transation
// param due - A deadline in minutes
// param version - A network version
// param network - A network id
// return - A common transaction struct
func CommonPart(txtype, version int, timeStamp, due int64, senderPublicKey string) base.CommonTransaction {
	return base.CommonTransaction{
		Type:      txtype,
		Version:   version,
		Signer:    senderPublicKey,
		TimeStamp: timeStamp,
		Deadline:  timeStamp + due*60,
	}
}
