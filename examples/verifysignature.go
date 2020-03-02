package main

import (
	"fmt"

	"github.com/isarq/nem-sdk-go/model"
	"github.com/isarq/nem-sdk-go/utils"
)

func main() {
	// Create keypair
	kp, _ := model.KeyPairCreate("")

	// Data to sign
	data := "NEM is awesome !"

	// Sign data
	sig, _ := kp.Sign([]byte(data))

	// Review
	fmt.Println("Public key: ", kp.PublicString())
	fmt.Println("Original data: ", data)
	fmt.Println("Signature: ", utils.Bt2Hex(sig))

	// Result
	fmt.Println()
	if model.Verify(kp.Public, []byte(data), sig) {
		fmt.Println("Signature is valid")
	} else {
		fmt.Println("Signature is invalid")
	}
}
