// 06_gomaxprocs
package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(1) //指定以单核运算
	fmt.Println("n = ", n)
	n = runtime.GOMAXPROCS(2) //指定以单核运算
	fmt.Println("n = ", n)
	for {
		go fmt.Print(1)
		fmt.Print(0)
	}
	fmt.Println("Hello World!")
}
