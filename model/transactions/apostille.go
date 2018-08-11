package transactions

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"github.com/isarq/nem-sdk-go/base"
	"github.com/isarq/nem-sdk-go/model"
	"github.com/isarq/nem-sdk-go/utils"
	"golang.org/x/crypto/sha3"
	"strings"
)

type Dedicated struct {
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
}

type Apost struct {
	name          string
	signedVersion string
	version       string
}

type Apostilledata struct {
	Data        DataAp      `json:"data"`
	Transaction base.TxDict `json:"transaction"`
}

type DataAp struct {
	File             File       `json:"file"`
	Hash             string     `json:"hash"`
	Checksum         string     `json:"checksum"`
	DedicatedAccount *Dedicated `json:"dedicatedAccount"`
	Tags             string     `json:"tags"`
}

type File struct {
	Name    string `json:"name"`
	Hash    string `json:"hash"`
	Content []byte `json:"content"`
}

// Apostille hashing methods with version bytes
var Hashing = map[string]Apost{
	"MD5":      {"MD5", "81", "01"},
	"SHA1":     {"SHA1", "82", "02"},
	"SHA256":   {"SHA256", "83", "03"},
	"SHA3-256": {"SHA3-256", "88", "08"},
	"SHA3-512": {"SHA3-512", "89", "09"},
}

// Hash the file content depending of hashing
// param data - File content
// param hashing - The chosen hashing object
// param isPrivate - True if apostille is private, false otherwise
// return - The file hash with checksum
func hashFileData(data []byte, hashing Apost, isPrivate bool) string {
	// Full checksum is 0xFE (added automatically if hex txes) + 0x4E + 0x54 + 0x59 + hashing version byte
	var checksum string
	if isPrivate {
		checksum = "4e5459" + hashing.signedVersion
	} else {
		checksum = "4e5459" + hashing.version
	}
	// Build the apostille hash
	switch hashing.name {
	case "MD5":
		hasher := md5.New()
		hasher.Write([]byte(data))
		return checksum + utils.Bt2Hex(hasher.Sum(nil))
	case "SHA1":
		hasher := sha1.New()
		hasher.Write([]byte(data))
		return checksum + utils.Bt2Hex(hasher.Sum(nil))
	case "SHA256":
		hasher := sha256.New()
		hasher.Write([]byte(data))
		return checksum + utils.Bt2Hex(hasher.Sum(nil))
	case "SHA3-256":
		hasher := sha3.New256()
		hasher.Write([]byte(data))
		return checksum + utils.Bt2Hex(hasher.Sum(nil))
	default:
		hasher := sha3.New512()
		hasher.Write([]byte(data))
		return checksum + utils.Bt2Hex(hasher.Sum(nil))
	}
}

// Create an apostille object
// param common - A common object
// param fileName - The file name (with extension)
// param fileContent - The file content
// param tags - The apostille tags
// param hashing - An hashing object
// param isMultisig - True if transaction is multisig, false otherwise
// param multisigAccount - An [AccountInfo] struct
// link https://bob.nem.ninja/docs/#accountInfo} object
// param isPrivate - True if apostille is private / transferable / updateable, false if public
// param network - A network id
// return - An apostille object containing apostille data and the prepared transaction ready to be sent
func Create(common Common, fileName string, fileContent []byte, tags string, hashing Apost, isMultisig bool,
	multisigAccount string, isPrivate bool, network int) Apostilledata {
	var dedicatedAccount Dedicated
	var apostilleHash string
	if isPrivate {
		// Create user keypair
		kp := model.KeyPairCreate(common.PrivateKey)
		// Create the dedicated account
		dedicatedAccount = generateAccount(common, fileName, network)

		// Create hash from file content and selected hashing
		hash := hashFileData(fileContent, hashing, isPrivate)

		// Get checksum
		checksum := hash[:8]

		// Get hash without checksum
		dataHash := hash[8:]

		// Set checksum + signed hash as message
		apostilleHash = checksum + utils.Bt2Hex(kp.Sign(dataHash))

	} else {
		dedicatedAccount.Address = strings.ToUpper(strings.Replace(model.Apostille[network], "-", "", -1))
		dedicatedAccount.PrivateKey = "None (public sink)"
		apostilleHash = hashFileData(fileContent, hashing, isPrivate)
	}

	// Create transfer transaction struct
	transaction := TransferA(dedicatedAccount.Address, 0, apostilleHash)
	// Multisig
	transaction.IsMultisig = isMultisig
	transaction.MultisigAccount = multisigAccount
	// Set message type to hexadecimal
	transaction.MessageType = 0
	// Prepare the transfer transaction object
	transactionEntity := transaction.Prepare(common, network)
	//fmt.Printf("%s", utils.Struc2Json(transactionEntity))

	return Apostilledata{
		Data: DataAp{
			File: File{
				Name:    fileName,
				Hash:    apostilleHash,
				Content: fileContent,
			},
			Hash:             "fe" + apostilleHash,
			Checksum:         "fe" + apostilleHash[0:8],
			DedicatedAccount: &dedicatedAccount,
			Tags:             tags,
		},
		Transaction: transactionEntity,
	}
}

// Verify an apostille
// param fileContent - The file content
// param apostilleTransaction - The transaction object for the apostille
// return - True if valid, false otherwise
func VerifyApost(fileContent []byte, apostilleTransaction base.TransactionResponce) bool {
	var apostilleHash string
	if apostilleTransaction.Type == 4100 {
		tx := apostilleTransaction.OtherTrans
		apostilleHash = tx.Message.Payload
	} else {
		apostilleHash = apostilleTransaction.Message.Payload
	}
	// Get the checksum
	var checksum = apostilleHash[:10]
	// Get the hashing byte
	var hashingByte = checksum[8:]

	// Retrieve the hashing method using the checksum in message and hash the file accordingly
	fileHash := retrieveHash(apostilleHash, fileContent)

	if isSigned(hashingByte) {
		pk := utils.Hex2Bt(apostilleTransaction.Signer)
		v := model.Verify(pk, []byte(fileHash), utils.Hex2Bt(apostilleHash[10:]))
		return v
	} else {
		// Check if hashed file match hash in transaction (without checksum)
		return fileHash == apostilleHash[10:]
	}
}

// Hash a file according to version byte in checksum
// param apostilleHash - The hash contained in the apostille transaction
// param fileContent - The file content
// return - The file content hashed with correct hashing method
func retrieveHash(apostilleHash string, fileContent []byte) string {
	// Get checksum
	checksum := apostilleHash[0:10]
	// Get the version byte
	hashingVersionBytes := string(checksum[8:])

	// Hash depending of version byte
	if hashingVersionBytes == "01" || hashingVersionBytes == "82" {
		hasher := md5.New()
		hasher.Write(fileContent)
		return utils.Bt2Hex(hasher.Sum(nil))
	} else if hashingVersionBytes == "02" || hashingVersionBytes == "82" {
		hasher := sha1.New()
		hasher.Write(fileContent)
		return utils.Bt2Hex(hasher.Sum(nil))
	} else if hashingVersionBytes == "03" || hashingVersionBytes == "83" {
		hasher := sha256.New()
		hasher.Write(fileContent)
		return utils.Bt2Hex(hasher.Sum(nil))
	} else if hashingVersionBytes == "08" || hashingVersionBytes == "88" {
		hasher := sha3.New256()
		hasher.Write(fileContent)
		return utils.Bt2Hex(hasher.Sum(nil))
	} else {
		hasher := sha3.New512()
		hasher.Write(fileContent)
		return utils.Bt2Hex(hasher.Sum(nil))
	}
}

// Check if an apostille is signed
// param hashingByte - An hashing version byte
// return - True if signed, false otherwise
func isSigned(hashingByte string) bool {
	for a := range Hashing {
		if Hashing[a].signedVersion == hashingByte {
			return true
		}
	}
	return false
}

// Generate the dedicated account for a file. It will always generate the same private key for a given file name and private key
// param common - A common object
// param fileName - The file name (with extension)
// param network - A network id
// return - An object containing address and private key of the dedicated account
func generateAccount(common Common, fileName string, network int) Dedicated {
	// Create user keypair
	kp := model.KeyPairCreate(common.PrivateKey)

	// Create recipient account from signed sha256 hash of new filename
	hasher := sha256.New()
	hasher.Write([]byte(fileName))
	signedFilename := kp.Sign(utils.Bt2Hex(hasher.Sum(nil)))

	// Truncate signed file name to get a 32 bytes private key
	dedicatedAccountPrivateKey := utils.FixPrivateKey(utils.Bt2Hex(signedFilename))

	// Create dedicated account key pair
	dedicatedAccountKeyPair := model.KeyPairCreate(common.PrivateKey)

	return Dedicated{
		Address:    model.ToAddress(dedicatedAccountKeyPair.PublicString(), network),
		PrivateKey: dedicatedAccountPrivateKey,
	}
}

// An un-prepared transfer transaction object
// param recipient - A NEM account address
// param amount - A number of XEM
// param message - A message
// return A - Transfer struct
func TransferA(recipient string, amount float64, message string) Transfer {
	return Transfer{
		Amount:             amount,
		Recipient:          recipient,
		RecipientPublicKey: "",
		IsMultisig:         false,
		Message:            message,
		MessageType:        1,
		Mosaics:            nil,
	}
}
