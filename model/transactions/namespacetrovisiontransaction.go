package transactions

import (
	"errors"
	"github.com/isarq/nem-sdk-go/base"
	"github.com/isarq/nem-sdk-go/extras"
	"github.com/isarq/nem-sdk-go/model"
	"github.com/isarq/nem-sdk-go/utils"

	"strings"
)

type nsPrepare struct {
	senderPublicKey string
	rentalFeeSink   string
	rentalFee       float64
	namespaceParent string
	namespaceName   string
	due             int64
	network         int
}

type NamespaceProvision struct {
	NamespaceName   string `json:"namespaceName,omitempty"`
	NamespaceParent struct {
		Fqn string `json:"fqn,omitempty"`
	} `json:"namespaceParent,omitempty"`
	IsMultisig      bool   `json:"isMultisig"`
	MultisigAccount string `json:"multisigAccount"`
}

func (r *NamespaceProvision) Get() {
	r.MultisigAccount = ""
	r.IsMultisig = false
}

func (t *NamespaceProvision) GetType() int {
	return 0
}

// Prepare a namespace provision transaction object
// param common - A common struct
// param r - An un-prepared namespaceProvisionTransaction method
// param network - A network id
// return - A [ProvisionNamespaceTransaction] struct
// link {http://bob.nem.ninja/docs/#provisionNamespaceTransaction}
func (r *NamespaceProvision) Prepare(common Common, network int) base.Transaction {
	var msc nsPrepare
	if extras.IsEmpty(common) || extras.IsEmpty(network) {
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

	msc.rentalFeeSink = strings.ToUpper(strings.Replace(model.Namespace[network], "-", "", -1))

	// Set fee depending if namespace or sub
	if !extras.IsEmpty(r.NamespaceParent) {
		msc.rentalFee = model.SubProvisionNamespaceTransaction
	} else {
		msc.rentalFee = model.RootProvisionNamespaceTransaction
	}

	msc.namespaceParent = r.NamespaceParent.Fqn

	msc.namespaceName = r.NamespaceName

	if network == model.Data.Testnet.ID {
		msc.due = 60
	} else {
		msc.due = 24 * 60
	}
	msc.network = network

	rt := construct(msc)
	if r.IsMultisig && r.MultisigAccount != "" {
		return MultisigWrapper(kp.PublicString(), rt, msc.due, network)
	}
	return rt
}

// Create a namespace provision transaction struct
// param msc network - A nsPrepare struct
// return - A [ProvisionNamespaceTransaction] struct
// link http://bob.nem.ninja/docs/#provisionNamespaceTransaction
func construct(msc nsPrepare) *base.ProvisionNamespaceTransaction {
	timeStamp := utils.CreateNEMTimeStamp()
	version := model.GetVersion(1, msc.network)
	data := CommonPart(model.ProvisionNamespace, version, timeStamp, msc.due, msc.senderPublicKey)
	fee := model.NamespaceAndMosaicCommon

	custom := base.ProvisionNamespaceTransaction{
		CommonTransaction: base.CommonTransaction{
			TimeStamp: data.TimeStamp,
			Signer:    data.Signer,
			Type:      data.Type,
			Deadline:  data.Deadline,
			Version:   data.Version,
			Fee:       fee,
		},
		RentalFeeSink: msc.rentalFeeSink,
		RentalFee:     msc.rentalFee,
		Parent:        msc.namespaceParent,
		NewPart:       msc.namespaceName,
	}
	return &custom
}
