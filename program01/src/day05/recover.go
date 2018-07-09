// recover
package main

import (
	"fmt"
)

func testa() {
	fmt.Println("aaaaaaaaaa")
}
func testb(x int) {
	//设置ｒｅｃｏｖｅｒ
	defer func() {
		//recover()
		//可以打印ｐａｎｉｃ　错误信息

		if err := recover(); err != nil {
			//产生了ｐａｎｉｃ 异常
			fmt.Println(err)
		}
	}()

	var a [10]int
	a[x] = 111
	fmt.Println("bbbbbbbbb")
}
func testc() {
	fmt.Println("ccccccccc")
}

func main() {
	testa()
	testb(20)
	testc()
	fmt.Println("Hello World!")
}
