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
	common := objects.GetCommon("", "265087519502bd6f6c93f74b189ecdea18da9f58ba9d83a425821e714ea2aeea", false)

	// Get a NamespaceProvisionTransaction struct
	tx := objects.Namespaceprovision()

	// Define the NameSpace
	// The new part which is concatenated to the parent with a '.' as separator.
	tx.NamespaceName = "toto2"
	// The NamespaceParent.Fqn. This can be null if the transaction rents a root namespace.
	//tx.NamespaceParent.Fqn = "tototo"

	// Enable Multisig
	//tx.IsMultisig = false

	// Publickey of the multifirm account (only if IsMultisig is true).
	//tx.MultisigAccount = "aef5822056f21d73790c5c82e9043a51d82b4a516bc3652e8e5385f56d9bc244"

	// Prepare the transaction struct
	transactionEntity := tx.Prepare(common, model.Data.Testnet.ID)

	res, err := transactions.Send(common, transactionEntity, client)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("NamespaceProvision\n%s", utils.Struc2Json(res))
}
