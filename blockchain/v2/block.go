package main

import (
	"time"
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

	//　初始化　区块　各个　字段的数据
	block := &Block{
		Version:1,
		PrevBlockHash:prevBlockHash,
		// Hash:
		TimeStamp:time.Now().Unix(),
		TargetBits:targetBits,
		Nonce:0,
		MerKelRoot:[]byte{},
		Data:[]byte(data),
	}
	//block.SetHash() // 设置　区块的哈希值
	pow := NewProofOfWork(block)
	nonce,hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash

	return block
}

/*
func (block *Block) SetHash(){
	tmp := [][]byte{
		// 实现　int 类型　转换为　byte 类型的工具函数
		IntToByte(block.Version),
		block.PrevBlockHash,
		IntToByte(block.TimeStamp),
		block.MerKelRoot,
		IntToByte(block.Nonce),
		block.Data,
	}
	//　将区块的 各个字段　链接成一个切片　使用[]byte{} 进行链接　目的是避免污染　原区块的信息
	data := bytes.Join(tmp,[]byte{})
	//　对区块进行　sha256哈希算法 返回值为　[32]byte　数组　　不是切片
	hash := sha256.Sum256(data)
	block.Hash = hash[:] //由数组转换成切片
}
*/
//　创建比特币　的创世块　　即第一个区块　　它的前一个区块的哈希值为空
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Blcok! ",[]byte{})
}
