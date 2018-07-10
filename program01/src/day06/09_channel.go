// 09_channel
package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个无缓存的channel
	ch := make(chan int, 0)

	// len(ch) 缓冲区　剩余数据个数　cap(ch)缓冲区大小
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	//新建一个　goroutine
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("子routine i = ", i)
			ch <- i //往chan 写东西
		}
	}()

	//延时
	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		num := <-ch
		//读管道中的内容，没有美容前　阻塞
		//放进去　没人读取，就会阻塞，等待读取之后，才能再写入
		fmt.Println(" num = ", num)
	}
	fmt.Println("Hello World!")
}
