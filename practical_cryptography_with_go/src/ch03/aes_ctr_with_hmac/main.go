package third

import "crypto/aes"
import (
	"../../../src"
	"errors"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
)

var (
	ErrEncrypt = errors.New("secret: encryption failed")
	ErrDecrypt = errors.New("secret: decryption failed")
)

// AES-CTR and AES-CBC with an HMAC.
// data is first encrypted with AES in the appropriate mode, then an HMAC is appended.
//  the nonce must be only ever used once with the same key. Reusing a nonce can be catastrophic

// AES-CTR
// which HMAC construction to use
// HMAC-SHA-256 for AES-128 and HMAC-SHA-384 for AES-256

// we’ll encrypt by selecting a random nonce, encrypting the data, and computing the MAC for the ciphertext.
// The nonce will be prepended to the message and the MAC appended.
// The message will be encrypted in-place.
// The key is expected to be the HMAC key appended to the AES key.

const (
	// The AES block size in bytes.
	NonceSize = aes.BlockSize
	// We’ll use HMAC-SHA-256 with AES-256.
	MACSize = 32 // Output size of HMAC-SHA-256
	CKeySize = 32 // Cipher key size - AES-256
	MKeySize = 32 // HMAC key size - HMAC-SHA-256
)

var KeySize = CKeySize + MKeySize

func Encrypt(key, message []byte) ([]byte, error) {
	if len(key) != KeySize {
		return nil, ErrEncrypt
	}

	nonce, err := util.RandBytes(NonceSize)
	if err != nil {
		return nil, ErrEncrypt
	}

	ct := make([]byte, len(message))

	// NewCipher only returns an error with an invalid key size , but the key size was checked at the beginning of the function.
	c, _ := aes.NewCipher(key[:CKeySize])

	ctr := cipher.NewCTR(c, nonce)
	ctr.XORKeyStream(ct, message)

	h := hmac.New(sha256.New, key[CKeySize:])
	ct = append(nonce, ct...)
	h.Write(ct)
	ct = h.Sum(ct)

	return ct, nil
}

//

func Decrypt(key, message []byte) ([]byte, error) {
	if len(key) != KeySize {
		return nil, ErrDecrypt
	}

	if len(message) <= (NonceSize + MACSize) {
		return nil, ErrDecrypt
	}

	macStart := len(message) - MACSize
	tag := message[macStart:]
	out := make([]byte, macStart-NonceSize)
	message = message[:macStart]

	h := hmac.New(sha256.New, key[CKeySize:])
	h.Write(message)
	mac := h.Sum(nil)
	if !hmac.Equal(mac, tag) {
		return nil, ErrDecrypt
	}

	c, _ := aes.NewCipher(key[:CKeySize])
	ctr := cipher.NewCTR(c, message[:NonceSize])
	ctr.XORKeyStream(out, message[NonceSize:])

	return out, nil
}

// TODO 测试
