// 17_testdefer.go
package main

import (
	"fmt"
)

func test(x int) {
	result := 100 / x
	fmt.Println("result = ", result)
}

func main() {
	fmt.Println("aaaaaaaaa")
	defer fmt.Println("11111111")
	fmt.Println("bbbbbbbbb")
	defer fmt.Println("2222222")
	//调用一个函数，导致内存出错
	test(0)
	defer test(0)

	fmt.Println("cccccccccc")
	defer fmt.Println("3333333")
	//ｄｅｆｅｒ先进后出
}
