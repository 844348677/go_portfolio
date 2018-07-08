// 11_random.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//设置种子
	//产生随机数
	//如果种子参数一样，产生随机数一样
	//rand.Seed(666)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		//随机很大的数
		//fmt.Println("rand = ", rand.Int())
		//随机一百以内
		fmt.Println("rand = ", rand.Intn(100))
	}

	var a [10]int
	n := len(a)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100)
		fmt.Printf("%d, ", a[i])
	}
	fmt.Println("\n")
	//冒泡排序，挨着的两个元素比较，升序　大于则交换
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d, ", a[i])
	}
	fmt.Println()
	fmt.Println("Hello World!")
}
