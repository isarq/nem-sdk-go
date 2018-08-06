package objects

import (
	"encoding/hex"
	. "github.com/isarq/nem-sdk-go/base"
	"github.com/isarq/nem-sdk-go/model/transactions"
)

// An endpoint object
// param host - A NIS uri
// param port - A port
// return
func Endpoint(host string, port int) (endpoint Node) {
	return Node{
		Host: host,
		Port: port,
	}
}

// A common object
// param password - A password
// param privateKey - A privateKey
// param isHW - True if hardware wallet, false otherwise
// return struct
func GetCommon(password, privateKey string, isHW bool) transactions.Common {
	return transactions.Common{
		Password:   password,
		PrivateKey: privateKey,
		IsHW:       isHW,
	}
}

// Contains message types with name
// return
func MessageTypes() (message []MessageType) {
	ms := MessageType{
		Value: 0,
		Name:  "Hexadecimal",
	}
	message = append(message, ms)
	ms = MessageType{
		Value: 1,
		Name:  "Unencrypted",
	}
	message = append(message, ms)
	ms = MessageType{
		Value: 2,
		Name:  "Encrypted",
	}
	message = append(message, ms)
	return
}

// A multisig cosignatory modification object
// param type - 1 if an addition, 2 if deletion
// param publicKey - An account public key
// return
func MultisigCosignatoryModification(tp int, publicKey string) ConsModif {
	a := ConsModif{
		ModificationType:   tp,
		CosignatoryAccount: publicKey,
	}
	return a
}

type ToString struct {
	Objet []byte
}

func (b *ToString) ToString() string {
	return hex.EncodeToString(b.Objet)
}
