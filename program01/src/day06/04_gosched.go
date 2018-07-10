// 04_gosched
package main

import (
	"fmt"
	"runtime"
)

func main() {

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go")
		}
	}()

	for i := 0; i < 2; i++ {
		//让出时间片　先让别的goroutine执行　它执行完了，再回来执行
		runtime.Gosched()
		fmt.Println("hello")

	}
	//直接结束了

	fmt.Println("Hello World!")
}
