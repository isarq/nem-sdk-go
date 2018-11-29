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

	// Create a common object holding key
	common := objects.GetCommon("", "056862b3dffbfd67a78172cf04c6a917325f2325f40cd48eea736f40b8a96d58", false)

	// Create an un-prepared transfer transaction
	tx := objects.Transfer("TD2YSVI5L2OKSLAPJBWN7XXFBKYCHVXMXY42GS64", 1, "Hello")

	//// Enable Multisig
	//tx.IsMultisig = true
	//
	// Publickey of the multifirm account (only if IsMultisig is true).
	//tx.MultisigAccount = "31efa466d2c0aee147397ec3bbe16354fd6fc10eb6710014c8d9a8924ad9b152"

	// Prepare the transfer transaction
	transactionEntity := tx.Prepare(common, model.Data.Testnet.ID)

	res, err := transactions.Send(common, transactionEntity, *client)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("Transfer\n%s", utils.Struc2Json(res))
}
