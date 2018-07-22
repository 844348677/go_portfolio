package main

import (
	"os"
	"fmt"
	"flag"
)

const Usage = `
	createchain -address ADDRESS "create block chain"
	./block addBlock --data DATA	"add a block to block chain"
	./block printchain           	"print all blocks"
`
type CLI struct {
	//bc *BlockChain
}
func (cli *CLI) Run(){
	if len(os.Args) <2 {
		fmt.Println("too few parameters! ",Usage)
		os.Exit(1)
	}
	addBlockCmd := flag.NewFlagSet("addBlock",flag.ExitOnError)
	createChainCmd := flag.NewFlagSet("createchain",flag.ExitOnError)
	printCmd := flag.NewFlagSet("printchain",flag.ExitOnError)

	addBlockCmdPara := addBlockCmd.String("data","","block info")
	createChainCmdPara := createChainCmd.String("address","","address info")

	switch os.Args[1] {
	case "createchain":
		err := createChainCmd.Parse(os.Args[2:])
		CheckErr("createChainCmd",err)
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		CheckErr("addBlock ",err)
		if addBlockCmd.Parsed(){
			if *addBlockCmdPara == ""{
				fmt.Println("data is empty , pls check !")
				os.Exit(1)
			}
			//cli.bc.AddBlock(*addBlockCmdPara)
			//cli.AddBlock(*addBlockCmdPara)

		}

	case "printchain":
		err := printCmd.Parse(os.Args[2:])
		CheckErr("printchain ",err)
		if printCmd.Parsed(){
			cli.PrintChain()
		}
	default:
		fmt.Println("invalid cmd \n",Usage)
	}

	if createChainCmd.Parsed(){
		if *createChainCmdPara== ""{
			fmt.Println("data is empty , pls check !")
			os.Exit(1)
		}
		//cli.bc.AddBlock(*addBlockCmdPara)
		cli.CreateChain(*createChainCmdPara)

	}
	
}