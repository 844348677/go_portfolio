// 05_goexit
package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("cccccccc")
	// 终止函数
	//return
	//终止所在的　goroutine
	runtime.Goexit()
	fmt.Println("dddddddddddd")
}

func main() {

	go func() {
		fmt.Println("aaaaaaaaaaaaaaaa")
		//调用了别的函数
		test()
		fmt.Println("bbbbbbbbbbbbbbb")
	}()
	//死循环，
	for {

	}
	fmt.Println("Hello World!")
}
