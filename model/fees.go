package model

import (
	"github.com/isarq/nem-sdk-go/base"
	"github.com/isarq/nem-sdk-go/extras"
	"math"
)

// The Fee structure's base fee
// type int
const baseTransactionFee = 3

// The Fee structure's Fee factor
// type float64
const CurrentFeeFactor = 0.05

// The multisignature transaction fee
// type float64
var Multisigtransaction = math.Floor((baseTransactionFee * CurrentFeeFactor) * 1000000)

// The provision namespace transaction rental fee for root namespace
// type float64
const RootProvisionNamespaceTransaction = 100 * 1000000

// The provision namespace transaction rental fee for sub-namespace
// type float64
const SubProvisionNamespaceTransaction = 10 * 1000000

// The mosaic definition transaction fee
// type float64
const MosaicDefinitionTransaction = 10 * 1000000

// The common transaction fee for namespaces and mosaics
// type float64
var NamespaceAndMosaicCommon = math.Floor((baseTransactionFee * CurrentFeeFactor) * 1000000)

// The cosignature transaction fee
// type float64
var SignatureTransaction = math.Floor((baseTransactionFee * CurrentFeeFactor) * 1000000)

// The importance transfer transaction fee
// type float64
var ImportanceTransferTransaction = math.Floor((baseTransactionFee * CurrentFeeFactor) * 1000000)

// The multisignature aggregate modification transaction fee
// type float64
var MultisigAggregateModificationTransaction = math.Floor((10 * CurrentFeeFactor) * 1000000)

// Calculate message fee. 0.05 XEM per commenced 32 bytes
// If the message is empty, the fee will be 0
// param message - An message struct
// param isHW - True if hardware wallet, false otherwise
// return - The message fee
func CalculateMessage(message base.Message, isHW bool) float64 {

	if extras.IsEmpty(message.Payload) {
		return 0.00
	}

	length := float64(len(message.Payload)/32 + 1)

	// Add salt and IV and round up to AES block size
	if isHW && message.Type == 2 {
		length = 32 + 16 + math.Ceil(float64(length/16))*16
	}
	return CurrentFeeFactor * math.Floor(length)
}

// Calculate fees from an amount of XEM
// param numNem - An amount of XEM
// return - The minimum fee
func CalculateMinimum(numNem float64) float64 {
	fee := math.Floor(math.Max(1, numNem/10000))
	if fee > 25 {
		return 25
	}
	return fee
}

// Calculate mosaic quantity equivalent in XEM
// param multiplier - A mosaic multiplier
// param q - A mosaic quantity
// param sup - A mosaic supply
// param divisibility - A mosaic divisibility
// return - The XEM equivalent of a mosaic quantity
func CalculateXemEquivalent(multiplier, q, sup, divisibility float64) float64 {
	if sup == 0 {
		return 0
	}
	// TODO: can this go out of JS (2^54) bounds? (possible BUG)
	return 8999999999 * q * multiplier / sup / math.Pow(10, divisibility+6)
}
