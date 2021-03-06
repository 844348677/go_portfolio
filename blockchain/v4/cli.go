package main

import (
	"os"
	"fmt"
	"flag"
)

const Usage = `
	createchain -address ADDRESS "create block chain"
	//
	./block addBlock --data DATA	"add a block to block chain"
	./block printchain           	"print all blocks"
	getbalance -address ADDRESSS "get balance"
	send -from FROM -to TO -amount AMOUNT "make a transaction"
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
	getbalanceCmd := flag.NewFlagSet("getbalance",flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send",flag.ExitOnError)

	addBlockCmdPara := addBlockCmd.String("data","","block info")
	createChainCmdPara := createChainCmd.String("address","","address info")
	getbalanceCmdPara := getbalanceCmd.String("address","","address info")

	fromPara := sendCmd.String("from","","from address info")
	toPara := sendCmd.String("to","","to address info")
	amountPara := sendCmd.Float64("amount",0,"amount value")

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
	case "getbalance":
		err := getbalanceCmd.Parse(os.Args[2:])
		CheckErr("getbalanceCmd ",err)
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		CheckErr("sendCmd ",err)
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
	if getbalanceCmd.Parsed(){
		if *getbalanceCmdPara == ""{
			fmt.Println(Usage)
			os.Exit(1)
		}
		cli.GetBalance(*getbalanceCmdPara)
	}

	if sendCmd.Parsed(){
		if *fromPara == "" || *toPara == "" || *amountPara ==0{
			fmt.Println(Usage)
			os.Exit(1)
		}
		cli.Send(*fromPara,*toPara,*amountPara)
	}
}