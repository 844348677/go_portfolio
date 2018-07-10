// 07_resource
package main

import (
	"fmt"
	"time"
)

//定义一个打印机　参数为字符串　按每个字符打印
//打印机属于公共资源
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}
func person1() {
	Printer("hello")
}
func person2() {
	Printer("world")
}
func main() {
	Printer("hello")
	Printer("world")

	//新建两个　goroutine　代表两个人　２个人同时使用打印机
	go person1()
	go person2()
	//特地不让主routine结束
	for {
	}
	fmt.Println("Hello World!")
}
