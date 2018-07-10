// 16_timer
package main

import (
	"fmt"
	"time"
)

func main() {
	//延时　２ｓ　打印一句话
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("时间到")

	time.Sleep(2 * time.Second)
	fmt.Println("时间到")

	<-time.After(2 * time.Second)
	//定时２ｓ　阻塞２ｓ后产生一个时间　往ｃｈａｎｎｅｌ写内容
	fmt.Println("Hello World!")
}
