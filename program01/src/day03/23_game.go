// 23_game.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateNum(p *int) {
	//设置种子
	rand.Seed(time.Now().UnixNano())
	var num int
	for {
		num = rand.Intn(10000)
		if num >= 1000 {
			break
		}
	}

	//fmt.Println("num = ", num)
	*p = num
}
func GetNum(s []int, num int) {
	s[0] = num / 1000 //取千位
	s[1] = (num % 1000) / 100
	s[2] = (num % 100) / 10
	s[3] = num % 10
}
func OnGame(randSlice []int) {
	var num int
	keySlice := make([]int, 4)
	for {
		for {
			fmt.Printf("请输入一个４位数: ")
			fmt.Scan(&num)
			if 999 < num && num < 10000 {
				break
			}
			fmt.Println("输入的数不符合要求")
		}
		fmt.Println("num = ", num)

		GetNum(keySlice, num)
		fmt.Println("keySlicke = ", keySlice)

		n := 0
		for i := 0; i < 4; i++ {
			if keySlice[i] > randSlice[i] {
				fmt.Printf("第 %d 位大了一点\n", i+1)
			} else if keySlice[i] < randSlice[i] {
				fmt.Printf("第 %d 位小了一点\n", i+1)
			} else {
				fmt.Printf("第 %d 位猜对了\n", i+1)
				n++
			}
		}
		if n == 4 {
			fmt.Println("全部猜对！！")
			break
		}
	}

}

func main() {
	var randNum int

	//产生一个４位随机数
	CreateNum(&randNum)
	fmt.Println(" randnum = ", randNum)
	randSlice := make([]int, 4)
	GetNum(randSlice, randNum)
	fmt.Println("randSlice = ", randSlice)

	n1 := 1234 / 1000
	n2 := (1234 % 1000) / 100
	n3 := (1234 % 100) / 10
	n4 := (1234 % 10)
	fmt.Printf("n1 = %d , n2 = %d, n3 = %d , n4 = %d \n", n1, n2, n3, n4)

	OnGame(randSlice)

}
