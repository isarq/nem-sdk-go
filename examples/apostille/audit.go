package main

import (
	"github.com/isarq/nem-sdk-go/com/requests"
	"github.com/isarq/nem-sdk-go/model"
	"github.com/isarq/nem-sdk-go/model/objects"
	"github.com/isarq/nem-sdk-go/utils"

	"fmt"
	"github.com/isarq/nem-sdk-go/model/transactions"
)

func main() {
	// Create an NIS endpoint
	endpoint := objects.Endpoint(model.DefaultTestnet, model.DefaultPort)
	client := requests.NewClient(endpoint)

	// Simulate the file content
	fileContent := []byte("Apostille is awesome !")

	// Just pass the file name
	//fileContent, err := ioutil.ReadFile("file.txt")
	//if err != nil {
	//	fmt.Print(err)
	//}

	// Transaction hash of the Apostille
	txHash := "3369f0f3b60d40f8083102409cb53a47856e078907c3bca1c7220ac0266f9722"

	rest, err := client.ByHash(txHash)
	if err != nil {
		fmt.Printf("Account data:\n%s", utils.Struc2Json(err))
		return
	}

	fmt.Printf("%s", utils.Struc2Json(rest))
	// Verify
	if transactions.VerifyApost(fileContent, rest.Transaction) {
		fmt.Println("Apostille is valid")
	} else {
		fmt.Println("Apostille is invalid")
	}
}
