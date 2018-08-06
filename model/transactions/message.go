package transactions

import (
	"github.com/isarq/nem-sdk-go/base"
	"github.com/isarq/nem-sdk-go/utils"
)

// Prepare a message struct
// param common - A common struct
// param tx - An un-prepared transferTransaction struct point
// return - A prepared message struct
func MsgPrepare(common Common, tx *Transfer) base.Message {
	if tx.MessageType == 2 && common.PrivateKey != "" {
		return base.Message{
			Type:    2,
			Payload: "",
		}
	} else if tx.MessageType == 2 && common.IsHW {
		return base.Message{
			Type:      2,
			Payload:   utils.Utf8ToHex(tx.Message),
			PublicKey: tx.RecipientPublicKey,
		}

	} else if tx.MessageType == 0 && utils.IsHexadecimal(tx.Message) {
		return base.Message{
			Type:    1,
			Payload: "fe" + tx.Message,
		}
	} else {
		return base.Message{
			Type:    1,
			Payload: utils.Utf8ToHex(tx.Message),
		}
	}
}
