// 18_defer.go
package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20
	defer fmt.Printf("defer：　a=%d , b = %d\n", a, b)

	a = 1
	b = 2
	defer func(a, b int) {
		fmt.Printf("a = %d,b=%d\n", a, b)
	}(a, b)
	//()代表调用此匿名函数
	//　提前把参数传递过去，只是还没有调用

	a = 111
	b = 222
	fmt.Printf("外部：　a=%d , b = %d\n", a, b)
}
