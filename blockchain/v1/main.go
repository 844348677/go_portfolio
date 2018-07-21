package main

import "fmt"

func main(){
	bc := NewBlockChain()

	bc.AddBlock("我转你一枚BTC")
	bc.AddBlock("我又转你一枚BTC")

	for i,block := range bc.blocks{
		fmt.Println(" --------- block num : ",i)
		fmt.Println("Version: ",block.Version)
		fmt.Printf("PrevBlockHash: %x\n",block.PrevBlockHash)
		fmt.Printf("Hash: %x\n",block.Hash)
		fmt.Printf("TimeStamp: %d\n",block.TimeStamp)
		fmt.Printf("MerKelRoot: %x\n",block.MerKelRoot)
		fmt.Printf("Nonce: %d\n",block.Nonce)

		fmt.Printf("Data: %s\n",block.Data)

	}
}
