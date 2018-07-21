package main

type BlockChain struct {
	blocks []*Block
}

func NewBlockChain() *BlockChain{
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func (bc *BlockChain) AddBlock(data string) {
	//最后一个区块的　hash
	lastBlock := bc.blocks[len(bc.blocks)-1]
	prevBlockHash := lastBlock.Hash
	block := NewBlock(data, prevBlockHash)
	bc.blocks = append(bc.blocks,block)
}