package crypto

import (
	"errors"
	nacl "github.com/ereyes01/cryptohelper"
	"github.com/isarq/nem-sdk-go/utils"
)

// Encode a message
// param senderPriv - A sender private key
// param recipientPub - A recipient public key
// param msg - A text message
// return Error - The encoded message
func Encode(senderPriv, recipientPub, msg string) (string, error) {
	// Errors
	if senderPriv == "" || recipientPub == "" || msg == "" {
		err := errors.New("Missing argument !")
		return "", err
	}
	if !utils.IsPrivateKeyValid(senderPriv) {
		err := errors.New("Private key is not valid !")
		return "", err
	}
	if !utils.IsPublicKeyValid(recipientPub) {
		err := errors.New("Public key is not valid !")
		return "", err
	}
	// Processing

	iv, _ := nacl.RandomKey(16)
	salt, _ := nacl.RandomKey(32)
	encoded, err := encode(senderPriv, recipientPub, msg, iv, salt)
	if err != nil {
		return "", err
	}
	// Result
	return encoded, nil
}

// Encode a message, separated from encode() to help testing
// param senderPriv - A sender private key
// param recipientPub - A recipient public key
// param msg - A text message
// param iv - An initialization vector
// param salt - A salt
// return - The encoded message
func encode(senderPriv, recipientPub, msg, iv, salt string) (string, error) {
	// Errors
	if senderPriv == "" || recipientPub == "" || msg == "" || iv == "" || salt == "" {
		err := errors.New("Missing argument !")
		return "", err
	}
	if !utils.IsPrivateKeyValid(senderPriv) {
		err := errors.New("Private key is not valid !")
		return "", err
	}
	if !utils.IsPublicKeyValid(recipientPub) {
		err := errors.New("Public key is not valid !")
		return "", err
	}
	// Processing
	//sk := utils.Hex2BaReversed(senderPriv)
	//pk := utils.Hex2Bt(recipientPub)
	//var shared [32]byte

	//r := key_derive(shared, salt, sk, pk)
	//encKey := r
	//encIv := {
	//    iv: convert.ua2words(iv, 16)
	//}
	//encrypted := CryptoJS.AES.encrypt(CryptoJS.enc.Hex.parse(convert.utf8ToHex(msg)), encKey, encIv)
	//// Result
	//result := convert.ua2hex(salt) + convert.ua2hex(iv) + CryptoJS.enc.Hex.stringify(encrypted.ciphertext)
	return "", nil
}

func keyDerive(shared [32]byte, salt, sk, pk string) {
	//nacl.lowlevel.crypto_shared_key_hash(shared, pk, sk, hashfunc);
	//for (let i = 0; i < salt.length; i++) {
	//    shared[i] ^= salt[i];
	//}
	//let hash = CryptoJS.SHA3(convert.ua2words(shared, 32), {
	//    outputLength: 256
	//});
	//return hash;
}
