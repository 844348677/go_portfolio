// panic
package main

import (
	"fmt"
)

func testa() {
	fmt.Println("aaaaaaaa")
}
func testb(x int) {
	fmt.Println("bbbbbbbbb")
	//显示调用ｐａｎｉｃ函数　导致程序中断
	//panic("this is a panic test b")
	var a [10]int
	a[x] = 111
	//当ｘ为２０的时候　导致数组越界　产生一个ｐａｎｉｃ　导致程序崩溃
	// index out of range
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
