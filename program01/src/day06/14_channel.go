// 14_channel
package main

import (
	"fmt"
)

//此通道　只能写　不能读
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

//此　channel 只能读　不能写
func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("num = ", num)
	}
}
func main() {
	//创建一个双向通道
	ch := make(chan int)

	//生产者　生产数字　写入channel
	//新开一个goroutine
	go producer(ch) //channel 传参　　引用传递

	//消费者，从channel读取内容　打印
	consumer(ch)

	fmt.Println("Hello World!")
}
