// 19_fibonacci
package main

import (
	"fmt"
	"time"
)

//ch只写　quit只读
func fibonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		//监听　ｃｈａｎｎｅｌ　数据流动
		select {
		case ch <- x:
			x, y = y, x+y
			fmt.Printf(" x = %d , y = %d \n", x, y)
		case flag := <-quit:
			fmt.Println("flag = ", flag)
			return
		}
	}
}

func main() {
	ch := make(chan int)    // 数字通信
	quit := make(chan bool) //程序是否结束

	//消费者　从channel 读取内容
	//新建　goroutine
	go func() {
		for i := 0; i < 8; i++ {
			num := <-ch
			//测试　停止１秒
			time.Sleep(1 * time.Second)
			fmt.Println(num)
		}
		//可以停止
		quit <- true
	}()

	//生产者　产生数字　写入channel
	fibonacci(ch, quit)
	fmt.Println("Hello World!")
}
