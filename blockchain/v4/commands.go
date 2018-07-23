package main

import (
	"fmt"
)

func (cli *CLI) CreateChain(address string){
	bc := NewBlockChain(address)
	defer bc.db.Close()
	fmt.Println("create blockchain successfully!")
}


/*
func (cli *CLI) AddBlock(data string){
	cli.bc.AddBlock(data)
	fmt.Println("AddBlock Successed")
}

*/

func (cli *CLI) PrintChain(){
	bc := GetBlockChainHandler()
	it := bc.Interator()
	for{
		block := it.Next() //取回当前hash指向的block 将hash值前移

		fmt.Println("Version: ",block.Version)
		fmt.Printf("PrevBlockHash: %x\n",block.PrevBlockHash)
		fmt.Printf("Hash: %x\n",block.Hash)
		fmt.Printf("TimeStamp: %d\n",block.TimeStamp)
		fmt.Printf("MerKelRoot: %x\n",block.MerKelRoot)
		fmt.Printf("Nonce: %d\n",block.Nonce)

		fmt.Printf("Transactions: %s\n",block.Transactions)
		pow := NewProofOfWork(block)
		fmt.Printf(" Isvalid :%v\n",pow.IsValid())

		if len(block.PrevBlockHash) == 0{
			break
		}
	}
}
func (cli *CLI) GetBalance(address string){
	bc := GetBlockChainHandler()
	defer bc.db.Close()

	var total float64
	utxos := bc.FindUTXOs(address)
	for _,utxos := range utxos{
		total += utxos.Value
	}
	fmt.Printf("The balance of '%s' is %f \n",address,total)
}
