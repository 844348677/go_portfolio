// 15_timer
package main

import (
	"fmt"
	"time"
)

//验证　time.NewTimer() 时间到了　只会响应一次
func main() {
	timer := time.NewTimer(1 * time.Second)

	for {
		<-timer.C
		fmt.Println("时间到")
	}
}

func main01() {
	//创建一个定时器　设置时间为２ｓ　2s后　往time通道写内容(当前时间)
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前时间: ", time.Now())

	// 2s 后　往　timer.C 写数据　有数据后　就可以读取
	t := <-timer.C //ｃｈａｎｎｅｌ　没有数据就会堵塞
	fmt.Println("t = ", t)
	fmt.Println("Hello World!")
}
