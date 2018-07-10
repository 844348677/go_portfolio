// 13_channel
package main

import (
	"fmt"
)

func main() {
	//死锁
	//不能运行
	//创建一个　channel 双向的
	ch := make(chan int)
	//双向　channel 能隐式转换为单向channel
	var writeCh chan<- int = ch //只能写不能读
	writeCh <- 666              //写
	// <-writeCh
	var readCh <-chan int = ch //只能读　不能写
	<-readCh
	// readCh<- 666

	//单向是无法转换为双向的
	//var ch2 chan int = writeCh
	fmt.Println("Hello World!")
}
