package main

import (
	"fmt"
	"github.com/isarq/nem-sdk-go/model"
)

func main() {

	keyPair, _ := model.KeyPairCreate("265087519502bd6f6c93f74b189ecdea18da9f58ba9d83a425821e714ea2aeea")

	publicKey := keyPair.PublicString()
	privateKey := keyPair.PrivateString()

	address, _ := model.ToAddress(publicKey, model.Data.Testnet.ID)

	fmt.Println("PrivateKey:\t", privateKey)

	fmt.Println("PublicKey:\t", publicKey)

	fmt.Print("Address:\t", address)
}
