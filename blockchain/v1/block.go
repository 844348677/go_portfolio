package main

import (
	"time"
	"crypto/sha256"
	"bytes"
)

type Block struct{
	Version int64
	//前一个区块的hash值
	PrevBlockHash []byte
	// 本区块的　hash　值　为了方便实现做的简化　正常　比特币不包含自己的hash
	Hash []byte
	//时间戳
	TimeStamp int64
	//难度值
	TargetBits int64
	Nonce int64
	MerKelRoot []byte

	Data []byte
}

func NewBlock(data string,prevBlockHash []byte) *Block{

	block := &Block{
		Version:1,
		PrevBlockHash:prevBlockHash,
		// Hash:
		TimeStamp:time.Now().Unix(),
		TargetBits:10,
		Nonce:5,
		MerKelRoot:[]byte{},
		Data:[]byte(data),
	}
	block.SetHash()
	return block
}

func (block *Block) SetHash(){
	tmp := [][]byte{
		IntToByte(block.Version),
		block.PrevBlockHash,
		IntToByte(block.TimeStamp),
		block.MerKelRoot,
		IntToByte(block.Nonce),
		block.Data,
	}
	data := bytes.Join(tmp,[]byte{})
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Blcok! ",[]byte{})
}