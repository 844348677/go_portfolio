package main

import (
	"os"
	"fmt"
	"flag"
)

const Usage = `
	./block addBlock --data DATA	"add a block to block chain"
	./block printchain           	"print all blocks"
`
type CLI struct {
	bc *BlockChain
}
func (cli *CLI) Run(){
	if len(os.Args) <2 {
		fmt.Println("too few parameters! ",Usage)
		os.Exit(1)
	}
	addBlockCmd := flag.NewFlagSet("addBlock",flag.ExitOnError)
	printCmd := flag.NewFlagSet("printchain",flag.ExitOnError)

	addBlockCmdPara := addBlockCmd.String("data","","block info")

	switch os.Args[1] {
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		CheckErr("addBlock ",err)
		if addBlockCmd.Parsed(){
			if *addBlockCmdPara == ""{
				fmt.Println("data is empty , pls check !")
				os.Exit(1)
			}
			//cli.bc.AddBlock(*addBlockCmdPara)
			cli.AddBlock(*addBlockCmdPara)

		}

	case "printchain":
		err := printCmd.Parse(os.Args[2:])
		CheckErr("printchain ",err)
		if printCmd.Parsed(){
			cli.PrintChain()
		}
	default:
		fmt.Println("invalid cmd ",Usage)
	}


	
}