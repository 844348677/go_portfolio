// 16_defer.go
package main

import (
	"fmt"
)

func main() {
	//defer　延迟调用　main函数结束前调用
	//　defer　先写后调用，都能执行到
	defer fmt.Println("cccccccc")
	fmt.Println("aaaaaaaaaaaa")

	//defer　延迟调用　main函数结束前调用
	defer fmt.Println("bbbbbbbbbbb")
}
