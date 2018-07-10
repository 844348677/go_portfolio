// 17_timer
package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	//重新设置
	timer.Reset(1 * time.Second)
	<-timer.C
	fmt.Println("时间到")
}

func main01() {
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("子 goroutine 可以打印了　，定时器时间到")
	}()
	//停止定时器
	timer.Stop()
	for {

	}

	fmt.Println("Hello World!")
}
