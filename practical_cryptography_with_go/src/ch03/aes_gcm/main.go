package main

import "errors"
import (
	"../../../src"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
)

// Advanced Encryption Standard
// If AES is required or chosen, AES-GCM is often the best choice
// it pairs the AES block cipher with the GCM block cipher mode.
// ??????
// It is an AEAD cipher: authenticated encryption with additional data.
// It encrypts some data, which will be authenticated along with some optional additional data that is not encrypted.
// The key length is 16 bytes for AES-128, 24 bytes for AES-192, or 32 bytes for AES-256.
// It also takes a nonce as input
// GCM is difficult to implement properly

// Cryptography Engineering ([Ferg10]) recommends using 256-bit keys

// The AEAD type in the crypto/cipher package uses the same “open” and “seal” terms as NaCl.

const (
	KeySize   = 32
	NonceSize = 12
)

// 新建　util 在src目录下

var (
	ErrEncrypt = errors.New("secret: encryption failed")
	ErrDecrypt = errors.New("secret: decryption failed")
)

// GenerateKey generates a new AES-256 key.
func GenerateKey() ([]byte, error) {
	return util.RandBytes(KeySize)
}

// GenerateNonce generates a new AES-GCM nonce.
func GenerateNonce() ([]byte, error) {
	return util.RandBytes(NonceSize)
}

// Encrypt secures a message using AES-GCM.
func Encrypt(key, message []byte) ([]byte, error) {
	// NewCipher creates and returns a new cipher.Block.
	// The key argument should be the AES key,
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, ErrEncrypt
	}
	// NewGCM returns the given 128-bit, block cipher wrapped in Galois Counter Mode with the standard nonce length.
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, ErrEncrypt
	}

	nonce, err := GenerateNonce()
	if err != nil {
		return nil, ErrEncrypt
	}

	// Seal will append the output to the first argument; the usage here appends the ciphertext to the nonce.
	// The final parameter is any additional data to be authenticated.
	out := gcm.Seal(nonce, nonce, message, nil)
	return out, nil

}

// This version does not provide any additional (unencrypted but authenticated) data in the ciphertext.

//the message is prefixed with a 32-bit sender ID, which allows the receiver to select the appropriate decryption key. The following example will authenticate this sender ID:

// EncryptWithID secures a message and prepends a 4-byte sender ID to the message.
func EncryptWithID(key, message []byte, sender uint32) ([]byte, error) {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, sender)

	// 代码是需要重构的
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, ErrEncrypt
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, ErrEncrypt
	}
	nonce, err := GenerateNonce()
	if err != nil {
		return nil, ErrEncrypt
	}
/*
	for n := range nonce{
		// 待测试
		buf = append(buf, byte(n))
	}
*/
	buf = append(buf, nonce...)
	out := gcm.Seal(buf, nonce, message, message[:4])
	return out, nil
}

var keyList = map[uint32][]byte{}
func SelectKeyForID(id uint32) ([]byte, bool) {
	k, ok := keyList[id]
	return k, ok
}

//
func DecryptWithID(message []byte) ([]byte, error) {
	if len(message) <= NonceSize+4 {
		return nil, ErrDecrypt
	}

	// SelectKeyForID is a mock call to a database or key cache.
	id := binary.BigEndian.Uint32(message[:4])
	key, ok := SelectKeyForID(id)
	if !ok {
		return nil, ErrDecrypt
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, ErrDecrypt
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, ErrDecrypt
	}

	nonce := make([]byte, NonceSize)
	copy(nonce, message[4:])

	// Decrypt the message, using the sender ID as the additional data requiring authentication.
	out, err := gcm.Open(nil, nonce, message[4+NonceSize:], message[:4])
	if err != nil {
		return nil, ErrDecrypt
	}
	return out, nil

}

// 待测试　decrypt


// Decrypt recovers a message secured using AES-GCM.
func Decrypt(key, message []byte) ([]byte, error) {
	if len(message) <= NonceSize {
		return nil, ErrDecrypt
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, ErrDecrypt
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, ErrDecrypt
	}

	nonce := make([]byte, NonceSize)
	copy(nonce, message)

	out, err := gcm.Open(nil, nonce, message[NonceSize:], nil)
	if err != nil {
		return nil, ErrDecrypt
	}
	return out, nil
}