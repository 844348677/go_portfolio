// 21_copy.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InitData(s []int) {
	//设置随机ｓｅｅｄ
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}
}
func BubbleSort(s []int) {
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func main() {
	srcSlice := []int{1, 2}
	dstSlice := []int{6, 6, 6, 6, 6}
	copy(dstSlice, srcSlice)
	//　把后面的ｓｌｉｃｅ　ｃｏｐｙ到前面目标
	fmt.Println("dst = ", dstSlice)

	n := 10
	s := make([]int, n)
	InitData(s) //初始化数据
	fmt.Println("s = ", s)

	BubbleSort(s) //冒泡排序
	fmt.Println("s = ", s)
}
