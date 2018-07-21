package main

import "os"

//　构造区块连　结构　使用数组来存储所有区块
type BlockChain struct {
	blocks []*Block
}

//　创建区块链实例　添加第一个创世块
func NewBlockChain() *BlockChain{
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

//　添加区块操作
func (bc *BlockChain) AddBlock(data string) {
	//校验　数组的元素个数　避免出现访问越界
	if len(bc.blocks) <=0{
		os.Exit(1)
	}
	//最后一个区块的　hash
	lastBlock := bc.blocks[len(bc.blocks)-1]
	prevBlockHash := lastBlock.Hash
	block := NewBlock(data, prevBlockHash)
	bc.blocks = append(bc.blocks,block)
}