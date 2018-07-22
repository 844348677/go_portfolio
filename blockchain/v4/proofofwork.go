package main

import (
	"math/big"
	"bytes"
	"math"
	"crypto/sha256"
	"fmt"
)

//难度值
const targetBits = 24

type ProofOfWork struct{
	block *Block
	targetBit *big.Int

}

func NewProofOfWork(block *Block) *ProofOfWork {
	var IntTarget = big.NewInt(1)

	//000000000000000000001
	//100000000000000000000 十进制
	//000000000000000100000　十进制
	//000000100000000000000 十六进制
	//000000000ad0sfsdf1121 实际值
	IntTarget.Lsh(IntTarget,uint(256 - targetBits))

	return &ProofOfWork{block,IntTarget}
}

func (pow *ProofOfWork) PrepareRawData(nonce int64) []byte{
	block := pow.block

	tmp := [][]byte{
		// 实现　int 类型　转换为　byte 类型的工具函数
		IntToByte(block.Version),
		block.PrevBlockHash,
		IntToByte(block.TimeStamp),
		block.MerKelRoot,
		IntToByte(nonce),
		IntToByte(targetBits),
		//block.Data,
		//block.Transactions, //TODO
	}
	//　将区块的 各个字段　链接成一个切片　使用[]byte{} 进行链接　目的是避免污染　原区块的信息
	data := bytes.Join(tmp,[]byte{})

	return data
}

func (pow *ProofOfWork) Run() (int64,[]byte) {
	var nonce int64
	var hash [32]byte
	var HashInt big.Int

	fmt.Println("begin mining... ")
	fmt.Printf("target hash : %x \n",pow.targetBit.Bytes())
	for nonce < math.MaxInt64{
		data := pow.PrepareRawData(nonce)
		hash = sha256.Sum256(data)

		HashInt.SetBytes(hash[:])

		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y

		//000000100000000000000 十六进制
		//000000000ad0sfsdf1121 实际值
		//　实际值　小于　target
		if HashInt.Cmp(pow.targetBit) == -1{
			fmt.Printf("Found Hash : %x\n",hash)
			break
		}else{
			nonce++
		}

		// if hash　对比　pow.targetBit
		//  hash 大于　target
		//  break
		// else
		// nonce ++

	}
	return nonce,hash[:]
}

func (pow *ProofOfWork) IsValid() bool{

	data := pow.PrepareRawData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	var IntHash big.Int
	IntHash.SetBytes(hash[:])
	return IntHash.Cmp(pow.targetBit) == -1
}