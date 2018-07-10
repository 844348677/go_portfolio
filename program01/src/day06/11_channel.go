// 11_channel
package main

import (
	"fmt"
	"time"
)

func main() {
	//创建　一个有缓存的ｃｈａｎｎｅｌ
	ch := make(chan int, 3)

	// len(ch) 缓冲区　剩余数据个数　cap(ch)缓冲区大小
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	go func() {
		for i := 0; i < 10; i++ {
			//fmt.Println("子routine i = ", i)
			ch <- i //往chan 写东西
			fmt.Printf("子协程[%d]: len(ch) = %d, cap(ch) = %d \n", i, len(ch), cap(ch))
		}
	}()

	time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {
		num := <-ch
		//读管道中的内容，没有美容前　阻塞
		//放进去　没人读取，就会阻塞，等待读取之后，才能再写入
		fmt.Println(" num = ", num)
	}
	//有缓存的时候就是异步的　发短信　缓存满了就是同步
	//没缓存的时候　同步　打电话

	fmt.Println("Hello World!")
}
