package main

import (
	"./bolt"
	"os"
	"fmt"
)

const dbfile = "blockChainDb.db"
const blockBucket = "block"
const lasthash = "lastHash"
const genesisBlockInfo = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks "

//　构造区块连　结构　使用数组来存储所有区块
type BlockChain struct {
	//blocks []*Block
	db *bolt.DB
	lastHash []byte
}

//　创建区块链实例　添加第一个创世块
func NewBlockChain(address string) *BlockChain{

	if IsBlockChainExist(){
		fmt.Println("Block Chain already exist! ")
		os.Exit(1)
	}

	//return &BlockChain{[]*Block{NewGenesisBlock()}}
	// func Open(path string, mode os.FileMode, options *Options) (*DB, error) {
	db,err := bolt.Open(dbfile,0600,nil)
	CheckErr("NewBlockChain1",err)

	var lastHash []byte
	// func (db *DB) Update(fn func(*Tx) error) error {
	db.Update(func(tx *bolt.Tx) error {

		//１　创建 bucket
		// 2 写数据
		coinbase := NewCoinbaseTX(address,genesisBlockInfo)
		genesis := NewGenesisBlock(coinbase)
		bucket,err := tx.CreateBucket([]byte(blockBucket))
		CheckErr("NewBlockChain2",err)

		err = bucket.Put(genesis.Hash,genesis.Serialize())
		CheckErr("NewBlockChain3",err)
		err = bucket.Put([]byte(lasthash),genesis.Hash)
		CheckErr("NewBlockChain4",err)
		lastHash = genesis.Hash

		return nil
	})

	return &BlockChain{db,lastHash}
}

//上面函数用来创建　下面函数用来返回
func GetBlockChainHandler() *BlockChain{
	if !IsBlockChainExist(){
		fmt.Println("Block Chain not exist, please create first! ")
		os.Exit(1)
	}

	db,err := bolt.Open(dbfile,0600,nil)
	CheckErr("GetBlockChainHandler 1",err)

	var lastHash []byte
	// func (db *DB) Update(fn func(*Tx) error) error {
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket != nil{
			// 数据不为空　读取　lastHash 即可
			lastHash = bucket.Get([]byte(lasthash))
		}
		return nil
	})
	CheckErr("GetBlockChainHandler 2",err)

	return &BlockChain{db,lastHash}
}

//　添加区块操作
//func (bc *BlockChain) AddBlock(data string) {
func (bc *BlockChain) AddBlock(transactions []*Transaction) {
	var prevBlockHash []byte
	err := bc.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		lasthash := bucket.Get([]byte(lasthash))
		prevBlockHash = lasthash
		return nil
	})
	CheckErr("AddBlock ",err)

	block := NewBlock(transactions,prevBlockHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))

		err := bucket.Put(block.Hash,block.Serialize())
		CheckErr("AddBlock2 ",err)
		err = bucket.Put([]byte(lasthash),block.Hash)
		CheckErr("AddBlock3 ",err)

		bc.lastHash = block.Hash

		return nil
	})
	CheckErr("AddBlock4 ",err)

}

type BlockChainIterator struct{
	db *bolt.DB
	currentHash []byte
}

func (bc *BlockChain) Interator() *BlockChainIterator{
	return &BlockChainIterator{bc.db,bc.lastHash}
}

func (it *BlockChainIterator) Next() *Block{
	var block *Block
	err := it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil{
			os.Exit(1)
		}
		blockTmp := bucket.Get(it.currentHash)
		block = Deserialize(blockTmp)
		it.currentHash = block.PrevBlockHash
		return nil
	})
	CheckErr("Next: ",err)
	return block
}

func IsBlockChainExist() bool{

	_,err := os.Stat(dbfile)
	if os.IsNotExist(err){
		return false
	}
	return true
}