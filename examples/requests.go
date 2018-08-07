package main

import (
	"github.com/isarq/nem-sdk-go/com/requests"
	"github.com/isarq/nem-sdk-go/model"
	"github.com/isarq/nem-sdk-go/utils"

	"fmt"
	"github.com/isarq/nem-sdk-go/model/objects"
)

func main() {

	// Address we'll use in some queries
	address := "TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S"
	address2 := "TD5OUIZXUYGWILTDDPLD64TK44HWFFQIPZRRXRIH"
	publickey := "0257b05f601ff829fdff84956fb5e3c65470a62375a1cc285779edd5ca3b42f6"

	// Create an NIS endpoint
	endpoint := objects.Endpoint(model.DefaultTestnet, model.DefaultPort)
	client := requests.NewClient(endpoint)

	// ******** ACCOUNT GETS ********

	// 1 - Get account data from account adress
	b, err := client.AccountData(address)
	if err != nil {
		fmt.Printf("Account data:\n%s", utils.Struc2Json(err))
		return
	}
	fmt.Printf("Account data:\n%s", utils.Struc2Json(b))

	// 2 - Get account data from public key
	c, err := client.AccountDataFromPublicKey(publickey)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("Account data:\n%s", utils.Struc2Json(c))

	// 3 - Gets the AccountMetaDataPair of an slice of accounts
	d, err := client.GetBatchAccountData([]string{address, address2})
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetBatchAccountData:\n%s", utils.Struc2Json(d))

	// ******** HISTORICAL GETS ********

	// 4 - Gets the AccountMetaDataPair of an account from a certain block.
	e, err := client.GetBatchHistoricalAccountData([]string{address, address2},
		104688770)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetBatchHistoricalAccountData:\n%s", utils.Struc2Json(e))

	// 5 - Gets the AccountMetaDataPair of an array of accounts from an historical height.
	f, err := client.GetHistoricalAccountData(address, 104688770)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("GetBatchHistoricalAccountData:\n%s", utils.Struc2Json(f))

	// ******** MOSAIC GETS ********

	// 6 - Gets an array of mosaic objects for a given account address.
	g, err := client.MosaicsOwned(address)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("MosaicsOwned:\n%s", utils.Struc2Json(g))

	// 7 - Gets an array of mosaic definition objects for a given account address.
	// The parent parameter is optional.
	h, err := client.MosaicDefinitionsCreated(address, "")
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("MosaicDefinitionsCreated:\n%s", utils.Struc2Json(h))

	// 8 - Get mosaic definitions of a namespace or sub-namespace.
	i, err := client.MosaicDefinitions("nw.fiat")
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("MosaicDefinitions:\n%s", utils.Struc2Json(i))

	// ******** NAMESPACE GETS ********

	// 9 - Gets an array of namespace objects for a given account address.
	// The parent parameter is optional. If supplied, only sub-namespaces of the parent namespace are returned.
	j, err := client.NamespacesOwned(address, "")
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("NamespacesOwned:\n%s", utils.Struc2Json(j))

	// ******** HARVESTING GETS ********

	// 10 - Get harvested blocks
	k, err := client.HarvestedBlocks(address)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("Harvested blocks:\n%s", utils.Struc2Json(k))

	// 11 - Starts harvesting
	err = client.StartHarvesting(publickey)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}

	// 12 - Stop harvesting
	err = client.StopHarvesting(publickey)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}

	// ******** TRANSACTIONS GETS ********

	// 13 - Get Incoming transactions
	l, err := client.IncomingTransactions(address, "", "")
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("Incoming transactions:\n%s", utils.Struc2Json(l))

	// 14 - Get Outgoing transactions
	n, err := client.OutgoingTransactions(address, "", "")
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("Outgoing transactions:\n%s", utils.Struc2Json(n))

	// 15 - Gets the array of transactions for which an account is the sender or receiver
	// and which have not yet been included in a block.
	m, err := client.UnconfirmedTransactions(address)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Println(utils.Struc2Json(m))

	// 16 - Gets all transactions of an account.
	o, err := client.AllTransactions(address, "", "")
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("Outgoing transactions:\n%s", utils.Struc2Json(o))

	// ******** VARIOUS GETS ********

	// 17 - Get chain height
	p, err := client.Height()
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("Harvested blocks:\n%s", utils.Struc2Json(p))

	// 18 - Get the current last block of the chain.
	q, err := client.LastBlock()
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("Harvested blocks:\n%s", utils.Struc2Json(q))

	// 19 - Get information about the maximum number of allowed harvesters and
	// how many harvesters are already using the node.
	r, err := client.UnlockInfo()
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("UnlockInfo:\n%s", utils.Struc2Json(r))

	// 20 - Gets the AccountMetaDataPair of the account for which the given account is the delegate account.
	s, err := client.Forwarded(address)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("Forwarded:\n%s", utils.Struc2Json(s))

}
