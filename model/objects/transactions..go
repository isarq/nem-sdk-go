package objects

import (
	"github.com/isarq/nem-sdk-go/base"
	"github.com/isarq/nem-sdk-go/model/transactions"
	"strings"
)

// An un-prepared transfer transaction object
// param recipient - A NEM account address
// param amount - A number of XEM
// param message - A message
// return A - Transfer struct
func Transfer(recipient string, amount float64, message string) transactions.Transfer {
	return transactions.Transfer{
		Amount:             amount,
		Recipient:          recipient,
		RecipientPublicKey: "",
		IsMultisig:         false,
		Message:            message,
		MessageType:        1,
		Mosaics:            nil,
	}
}

// An un-prepared signature transaction struct
// param multisigAccount - The multisig account address
// param txHash - The multisig transaction hash
// return A - SignatureT struct
func Signature(multisigAccount, txHash string) base.SignatureT {
	var compressedAccount string
	if len(multisigAccount) != 0 {
		compressedAccount = strings.ToUpper(strings.Replace(multisigAccount, "-", "", -1))
	}

	return base.SignatureT{
		OtherHash:    struct{ Data string }{Data: txHash},
		OtherAccount: compressedAccount,
	}
}

// An un-prepared mosaic definition transaction object
// return  A - MosaicDefinition struct
func Mosaicdefinition() (m *transactions.MosaicDefinition) {
	var tx transactions.MosaicDefinition
	tx.Get()
	return &tx
}

// An un-prepared mosaic supply change transaction object
// return A - Supply struct
func MosaicSupplyChange() base.Supply {
	return base.Supply{
		Mosaic:          "",
		SupplyType:      1,
		Delta:           0,
		IsMultisig:      false,
		MultisigAccount: "",
	}
}

// An un-prepared multisig aggregate modification transaction object
// return A - MultisigAggregateModific struct
func MultisigAggregateModification() base.MultisigAggregateModific {
	return base.MultisigAggregateModific{
		Modifications:   nil,
		RelativeChange:  nil,
		MultisigAccount: "",
		IsMultisig:      false,
	}
}

// An un-prepared namespace provision transaction object
// return A - NamespaceProvision
func Namespaceprovision() *transactions.NamespaceProvision {
	var tx transactions.NamespaceProvision
	tx.Get()
	return &tx
}

// An un-prepared importance transfer transaction object
// param remoteAccount - A remote public key
// param mode - 1 for activating, 2 for deactivating
// return A - ImportanceTransfer struct
func Importancetransfer(remoteAccount string, mode int) base.ImportanceTransfer {
	return base.ImportanceTransfer{
		RemoteAccount:   remoteAccount,
		Mode:            mode,
		MultisigAccount: "",
		IsMultisig:      false,
	}
}
