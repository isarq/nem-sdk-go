package transactions

import (
	"github.com/isarq/nem-sdk-go/base"
	"github.com/isarq/nem-sdk-go/model"
	"github.com/isarq/nem-sdk-go/utils"
)

// Wrap a transaction in a multisignature transaction
// param senderPublicKey - The sender public key
// param innerEntity - The transaction entity to wrap
// param due - The transaction deadline in minutes
// param network - A network id
// return - A [MultisigTransaction] struct
// link http://bob.nem.ninja/docs/#multisigTransaction
func MultisigWrapper(senderPublicKey string, innerEntity base.Tx, due int64, network int) *base.MultiSignTransaction {
	timeStamp := utils.CreateNEMTimeStamp()
	version := model.GetVersion(1, network)
	data := CommonPart(model.MultiSignTransaction, version, timeStamp, due, senderPublicKey)

	custom := base.MultiSignTransaction{
		CommonTransaction: base.CommonTransaction{
			TimeStamp: data.TimeStamp,
			Version:   data.Version,
			Signer:    data.Signer,
			Type:      data.Type,
			Deadline:  data.Deadline,
			Fee:       model.Multisigtransaction,
		},

		OtherTrans: innerEntity,
	}
	return &custom
}
