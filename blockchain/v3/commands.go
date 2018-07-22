package main

import "fmt"

func (cli *CLI) AddBlock(data string){
	cli.bc.AddBlock(data)
	fmt.Println("AddBlock Successed")
}
func (cli *CLI) PrintChain(){
	bc := cli.bc
	it := bc.Interator()
	for{
		block := it.Next() //取回当前hash指向的block 将hash值前移

		fmt.Println("Version: ",block.Version)
		fmt.Printf("PrevBlockHash: %x\n",block.PrevBlockHash)
		fmt.Printf("Hash: %x\n",block.Hash)
		fmt.Printf("TimeStamp: %d\n",block.TimeStamp)
		fmt.Printf("MerKelRoot: %x\n",block.MerKelRoot)
		fmt.Printf("Nonce: %d\n",block.Nonce)

		fmt.Printf("Data: %s\n",block.Data)
		pow := NewProofOfWork(block)
		fmt.Printf(" Isvalid :%v\n",pow.IsValid())

		if len(block.PrevBlockHash) == 0{
			break
		}
	}
}
