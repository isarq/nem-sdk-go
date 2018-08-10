package main

import (
	"fmt"
	"github.com/isarq/nem-sdk-go/com/requests"
	"github.com/isarq/nem-sdk-go/model"
	"github.com/isarq/nem-sdk-go/model/objects"
	"github.com/isarq/nem-sdk-go/model/transactions"
	"github.com/isarq/nem-sdk-go/utils"
)

func main() {
	// Create an NIS endpoint
	endpoint := objects.Endpoint(model.DefaultTestnet, model.DefaultPort)
	client := requests.NewClient(endpoint)

	// Create a common object holding key
	common := objects.GetCommon("", "064862b3dffbfd67a78172cf04c6a917325f2325f40cd48eea736f40b8b96d49", false)

	// Create variable to store our mosaic definitions,
	// needed to calculate fees properly (already contains xem definition)
	mosaicDefinitionMetaDataPair := objects.MosaicDefinitionMetadataPair()
	fmt.Sprint(mosaicDefinitionMetaDataPair)
	// Create an un-prepared mosaic transfer transaction struct
	// (use same object as transfer tansaction)
	tx := objects.Transfer("TCSBBN-7XUDLR-OZXZYJ-RCDZQC-33T3HE-FM3B4E-SESM", 1, "")

	//// Enable Multisig
	//tx.IsMultisig = true
	//
	//// Publickey of the multifirm account (only if IsMultisig is true).
	//tx.MultisigAccount = "31efa466d2c0aee147397ec3bbe16354fd6fc10eb6710014c8d9a8924ad9b152"

	// ATTACHING XEM MOSAIC
	// No need to get mosaic definition because it is already known in the mosaicdefinitionMetaDatapair

	// Create a mosaic attachment struct
	mosaicAttachment := objects.Attachment("nem", "xem", 5)

	// Append attachment into transaction mosaics
	tx.Mosaics = append(tx.Mosaics, mosaicAttachment)

	// ATTACHING ANOTHER MOSAIC
	// Need to get mosaic definition using com.requests
	// 100 nw.fiat.eur (divisibility is 2 for this mosaic)
	mosaicAttachment2 := objects.Attachment("ven", "ptr", 8)

	// Append attachment into transaction mosaics
	tx.Mosaics = append(tx.Mosaics, mosaicAttachment2)

	// Need mosaic definition of nw.fiat:eur to calculate adequate fees, so we get it from network.
	// Otherwise you can simply take the mosaic definition from
	// api manually (http://bob.nem.ninja/docs/#retrieving-mosaic-definitions)
	// and put it into mosaicDefinitionMetaDataPair model (objects.js)
	// next to nem:xem (be careful to respect object structure)
	mosaics, err := client.MosaicDefinitions("ven")
	// Check if the mosaic was found
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	//fmt.Println(utils.Struc2Json(mosaics))

	// Look for the mosaic definition(s) we want in the request response
	// (Could use []string{"eur", "usd"} to return eur and usd mosaicDefinitionMetaDataPairs)
	neededDefinition := requests.SearchMosaicDefinitionArray(mosaics, []string{"ptr"})

	// Get full name of mosaic to use as object key
	fullMosaicName := utils.MosaicIdToName(mosaicAttachment2.MosaicID)

	//// Set eur mosaic definition into mosaicDefinitionMetaDataPair
	mosaicDefinitionMetaDataPair[fullMosaicName] = neededDefinition[fullMosaicName]

	// Prepare the transfer transaction object
	transactionEntity := tx.PrepareMosaic(common, mosaicDefinitionMetaDataPair, client, model.Data.Testnet.ID)

	res, err := transactions.Send(common, transactionEntity, client)
	if err != nil {
		fmt.Println(utils.Struc2Json(err))
		return
	}
	fmt.Printf("MosaicDefinition:\n%s", utils.Struc2Json(res))

}
