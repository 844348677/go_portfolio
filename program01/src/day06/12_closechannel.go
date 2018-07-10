// 12_closechannel
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	//创建一个无缓存　channel
	//创建一个　goroutine
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //往通道数据
		}
		//不需要再写数据时　关闭channel
		close(ch)
		//关闭ｃｈａｎｎｅｌ不能再发数据　否则ｐａｎｉｃ
		//如果还有数据　可以继续读
	}() //别忘了使用()调用
	//range channel 当close channel 自动跳出循环
	for num := range ch {
		fmt.Println("num = ", num)
	}

	fmt.Println("Hello World!")
}
func main01() {
	ch := make(chan int)
	//创建一个无缓存　channel
	//创建一个　goroutine
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //往通道数据
		}
		//不需要再写数据时　关闭channel
		close(ch)
		//关闭ｃｈａｎｎｅｌ不能再发数据　否则ｐａｎｉｃ
		//如果还有数据　可以继续读
	}() //别忘了使用()调用
	for {
		//如果ｏｋ为ｔｒｕｅ 说明管道没有关闭
		if num, ok := <-ch; ok == true {
			fmt.Println("num = ", num)
		} else {
			//channel 关闭了　就跳出循环
			break
		}

	}
	fmt.Println("Hello World!")
}
