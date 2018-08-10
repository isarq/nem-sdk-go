package objects

import (
	"github.com/isarq/nem-sdk-go/base"
)

// A mosaic attachment object
// param namespaceId - A namespace name
// param mosaicName - A mosaic name
// param quantity - A mosaic quantity (in uXEM)
// return
func Attachment(namespaceId, mosaicName string, quantity float64) base.Mosaic {
	return base.Mosaic{
		MosaicID: base.MosaicID{
			NamespaceID: namespaceId,
			Name:        mosaicName,
		},
		Quantity: quantity,
	}
}

func MosaicDefinitionMetadataPair() map[string]base.MosaicDefinition {
	rest := make(map[string]base.MosaicDefinition)
	rest["nem:xem"] = base.MosaicDefinition{
		Creator:     "3e82e1c1e4a75adaa3cba8c101c3cd31d9817a2eb966eb3b511fb2ed45b8e262",
		Description: "xem",
		ID:          base.MosaicID{NamespaceID: "nem", Name: "xem"},
		Properties: []base.Properties{
			{Name: "divisibility", Value: "6"},
			{Name: "initialSupply", Value: "8999999999"},
			{Name: "supplyMutable", Value: "false"},
			{Name: "transferable", Value: "true"},
		},
	}
	return rest
}
