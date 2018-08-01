package main

import (
	"io"
	"crypto/rand"
	"fmt"
	"bytes"
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
