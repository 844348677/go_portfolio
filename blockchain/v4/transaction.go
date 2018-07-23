package main

import (
	"crypto/sha256"
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

const reword float64 = 12.5

type Transaction struct {
	TXID []byte
	TXInputs []Input
	TXOutputs []Output
}
type Input struct{
	Txid []byte
	ReferOutputIndex int64
	UnlockScript string
	// ScriptSig
}
type Output struct{
	Value float64
	LockScript string
	//ScriptPubKey
}

func (tx *Transaction) SetTXID(){
	//data := bytes.Join([][]byte{},[]byte{})
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	encoder.Encode(tx)

	hash := sha256.Sum256(buffer.Bytes())
	tx.TXID = hash[:]
}

func (input *Input) CanUnlockUTXOByAddress(unlockdata string) bool{
	// 简单　实现
	return input.UnlockScript == unlockdata
}

func (output *Output) CanBeUnlockedByAddress(unlockdata string) bool{
	return output.LockScript == unlockdata
}

func NewCoinbaseTX(address string, data string) *Transaction {

	if data == ""{
		data = fmt.Sprintf(data,"Current reward is : %f\n",reword)
	}
	input := Input{nil,-1,data}
	output := Output{reword,address}

	tx := Transaction{nil,[]Input{input},[]Output{output}}
	tx.SetTXID()
	return &tx
}

func (tx *Transaction) IsCoinbase() bool{
	// 使用元素下标　判断个数　满足下标
	if len(tx.TXInputs) == 1{
		if tx.TXInputs[0].Txid == nil && tx.TXInputs[0].ReferOutputIndex == -1{
			return true
		}
	}
	return false
}

func NewTransaction(from , to string , amount float64 , bc *BlockChain) *Transaction{
	counted, container := bc.FindSuitableUTXOs(from,amount)

	if counted < amount{
		fmt.Println("not enough founds ! ")
		os.Exit(1)
	}

	var inputs []Input
	var outputs []Output
	// 	var container = make(map[string][]int64)
	for txid, outputIndexs := range container{
		for _,index := range outputIndexs{
			input := Input{[]byte(txid),index,from}
			inputs = append(inputs,input)
		}
	}

	output := Output{amount,to}
	outputs = append(outputs,output)

	//多余的　钱给自己
	if counted > amount{
		outputs = append(outputs,Output{counted-amount,from})
	}

	tx := Transaction{nil,inputs,outputs}
	tx.SetTXID()
	return &tx
}