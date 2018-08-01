package main

import (
	"io"
	"crypto/rand"
	"fmt"
	"bytes"
	"errors"
	"golang.org/x/crypto/nacl/secretbox"
)

// It uses a 32-byte key and 24-byte nonces
const (

	KeySize = 32
	NonceSize = 24
)

// GenerateKey creates a new random secret key.
func GenerateKey() (*[KeySize]byte, error){
	// [KeySize]byte array
	key := new([KeySize]byte)
	// ReadFull reads exactly len(buf) bytes from r into buf.
	_, err := io.ReadFull(rand.Reader, key[:])

	// fmt.Println("key :" , key)

	if err != nil {
		return nil, err
	}
	return key, nil
}
// GenerateNonce creates a new random nonce.
func GenerateNonce() (*[NonceSize]byte, error) {
	nonce := new([NonceSize]byte)
	_, err := io.ReadFull(rand.Reader, nonce[:])

	if err != nil {
		return nil, err
	}
	return nonce, nil
}

// NaCl uses the term seal to mean securing a message (such that it now is confidential, and that its integrity and authenticity may be verified)
// and open to mean recovering a message (verifying its integrity and authenticity, and decrypting the message).

// 这里的写法够古老　文档有rand 例子
// https://golang.google.cn/pkg/crypto/rand/
func document_rand(){
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// The slice should now contain random bytes instead of only zeroes.
	fmt.Println(bytes.Equal(b, make([]byte, c)))
	fmt.Println("b : ",b)
}

// https://nacl.cr.yp.to/index.html

var (
	ErrEncrypt = errors.New("secret: encryption failed")
	ErrDecrypt = errors.New("secret: decryption failed")
)

// Encrypt generates a random nonce and encrypts the input using NaCl's secretbox package.
// The nonce is prepended to the ciphertext.
// A sealed message will the same size as the original message plus secretbox.Overhead bytes long.
func Encrypt(key *[KeySize]byte, message []byte) ([]byte, error) {
	// randomly generated nonces will be used;
	nonce, err := GenerateNonce()
	if err != nil {
		return nil, ErrEncrypt
	}

	// The recipient will need some way of recovering the nonce, so it will be prepended to the message.

	out := make([]byte, len(nonce))
	copy(out, nonce[:])
	fmt.Println("copy out : ",out)
	out = secretbox.Seal(out, message, nonce, key)
	return out, nil
}

// The length of messages is not hidden.


// Decryption expects that the message contains a prepended nonce, and we verify this assumption by checking the length of the message.
// A message that is too short to be a valid encryption message is dropped right away.





