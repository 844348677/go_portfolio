// 09_iota.go
package main

import (
	"fmt"
)

func main() {
	//　iota　常量自动生成器　每个一行　自动加一
	//　ｃｏｎｓｔ　赋值用
	const (
		a = iota
		b = iota
		c
	)
	fmt.Printf("a = %d, b=%d , c=%d \n", a, b, c)
	fmt.Println("Hello World!")
	// iota　重置为０
	const d = iota
	fmt.Println(d)
	//　如果同一行　值都一样
	const (
		i          = iota
		j1, j2, j3 = iota, iota, iota
		k          = iota
	)
	fmt.Printf("i  = %d, j1=%d , j2=%d ,j3=%d,k=%d \n", i, j1, j2, j3, k)
}
