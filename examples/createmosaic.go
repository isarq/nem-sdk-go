package main

import (
	"github.com/isarq/nem-sdk-go/com/requests"
	"github.com/isarq/nem-sdk-go/model"
	"github.com/isarq/nem-sdk-go/model/objects"
	"github.com/isarq/nem-sdk-go/utils"

	"fmt"
	"github.com/isarq/nem-sdk-go/base"
	"github.com/isarq/nem-sdk-go/model/transactions"
)

func main() {
	// Create an NIS endpoint
	endpoint := objects.Endpoint(model.DefaultTestnet, model.DefaultPort)
	client := requests.NewClient(endpoint)

	// Create a common object holding key
	common := objects.GetCommon("", "265087519502bd6f6c93f74b189ecdea18da9f58ba9d83a425821e714ea2aeea", false)

	// Get a MosaicDefinitionCreationTransaction struct
	tx := objects.Mosaicdefinition()

	// Enable Multisig
	//tx.IsMultisig = false

	// Publickey of the multifirm account (only if IsMultisig is true).
	//tx.MultisigAccount = "00244b414eefef48a34de44fbdf613aeb5925e2d652a101924c43c7f91f60e0e"

	// The MosaicName which is concatenated to the parent with a '.' as separator.
	tx.MosaicName = "nem-sdk-go"

	// The parent namespace.
	tx.NamespaceParent.Fqn = "isarq"

	tx.MosaicDescription = "My mosaic test from sdk Golang"

	// Set properties (see https://nemproject.github.io/#mosaicProperties)
	tx.Properties = []base.Properties{
		{Name: "divisibility", Value: "6"},
		{Name: "initialSupply", Value: "1000000000"},
		{Name: "transferable", Value: "true"},
		{Name: "supplyMutable", Value: "true"},
	}

	// Set Levy (see https://nemproject.github.io/#mosaicLevy)
	tx.Levy.FeeType = 0x01
	tx.Levy.Address = "TB3YJTWKY5IY62ABUIDLJ3YVEPX56OSVWULCQSWJ"
	tx.Levy.Mosaic.NamespaceID = "nem"
	tx.Levy.Mosaic.Name = "xem"
	tx.Levy.Fee = 400000

	transactionEntity := tx.Prepare(common, model.Data.Testnet.ID)

	res, err := transactions.Send(common, transactionEntity, client)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("MosaicDefinition:\n%s", utils.Struc2Json(res))
}
