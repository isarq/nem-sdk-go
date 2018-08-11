package main

import (
	"fmt"
	"github.com/isarq/nem-sdk-go/com/requests"
	"github.com/isarq/nem-sdk-go/model"
	"github.com/isarq/nem-sdk-go/model/objects"
	"github.com/isarq/nem-sdk-go/model/transactions"
	"github.com/isarq/nem-sdk-go/utils"
	"strings"
)

func main() {
	// Create an NIS endpoint
	endpoint := objects.Endpoint(model.DefaultTestnet, model.DefaultPort)
	client := requests.NewClient(endpoint)

	// Create a common object holding key
	common := objects.GetCommon("", "056862b3dffbfd67a78172cf04c6a917325f2325f40cd48eea736f40b8a96d58", false)

	// Enable Multisig
	IsMultisig := false
	// Publickey of the multifirm account (only if IsMultisig is true).
	MultiSignAccount := "31efa466d2c0aee147397ec3bbe16354fd6fc10eb6710014c8d9a8924ad9b152"

	// Simulate the file content
	fileContent := []byte("Apostille is awesome !")

	// Just pass the file name
	//fileContent, err := ioutil.ReadFile("file.txt")
	//if err != nil {
	//   fmt.Print(err)
	//}

	// Create the apostille
	apostille := transactions.Create(common, "file.txt", fileContent, "Test Apostille",
		transactions.Hashing["SHA256"], IsMultisig, MultiSignAccount, false, model.Data.Testnet.ID)

	// Serialize transfer transaction and announce
	res, err := transactions.Send(common, apostille.Transaction, client)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}

	fmt.Printf("%s", utils.Struc2Json(res))
	fmt.Println("Create a file with the fileContent text and name it: ", apostille.Data.File.Name)
	fmt.Println("-- Apostille TX ", res.TransactionHash.Data)
	fmt.Println("-- Date DD/MM/YYYY ", strings.Split(apostille.Data.File.Name, ".")[0])
	fmt.Println("When transaction is confirmed the file should audit successfully in Nano")
	fmt.Printf("You can also take the following hash: %s and put it into the audit.go example \n",
		res.TransactionHash.Data)
}
