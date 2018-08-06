# Nem-sdk-go 
[![N|Solid](http://isarq.com/descargas/Nem-go.png)](https://nodesource.com/products/nsolid)

NEM Developer Kit for Golang based on https://github.com/QuantumMechanics/NEM-sdk

# types of requests
### Account gets
  - Get account data from account address
  - Get account data from public key
  - Gets the AccountMetaDataPair of an array of accounts
### Transactions gets
- Get Incoming transactions.
- Get Outgoing transactions.
- Gets the array of transactions for which an account is the sender or receiver
	and which have not yet been included in a block.
- Gets all transactions of an account.

### Historical gets
  - Gets the AccountMetaDataPair of an account from a certain block.
  - Gets the AccountMetaDataPair of an array of accounts from an historical height.
### Mosaic gets
  - Gets an array of mosaic objects for a given account address.
  - Gets an array of mosaic definition objects for a given account address.
  - Gets mosaic definitions that an account owns.
### Namespace gets
  - Gets an array of namespace objects for a given account address.
### Harvesting gets
  - Get harvested blocks.
  - Starts harvesting.
  - Stop harvesting.
### Various gets
  - Get chain height.
  - Get the current last block of the chain.
  - Get information about the maximum number of allowed harvesters and
	how many harvesters are already using the node.
  - Gets the AccountMetaDataPair of the account for which the given 
    account is the delegate account.
 
# types of transactions!
  - Simple transactions.
  - Create mosaic.
  - Create namespace.
  - Multi-signature transactions.
  ### Other functions.
 - Create private keys.
 - Create key pairs.
 - Create key pairs.
 - Extract public key from key pair.
 - Verify a signature.
 - Convert public key to an address.
 - Verify address validity.
 - Verify if address is from given network.
 - More.
# features in development!
 - Mosaic transactions.
  - WebSocket

### Installation

```sh
$ go get -u github.com/isarq/nem-sdk-go
```

### Development

Want to contribute? Great!


### Todos

 - Write MORE Tests

License
----

MIT


**This project is in full development and many things can change!**
