// 05_Print
package main

import (
	"fmt"
)

func test() (a, b, c int) {
	return 1, 2, 3
}
func main() {
	a := 10
	fmt.Println("a = ", a)
	//格式化处理　输出
	fmt.Printf("a = %d\n", a)
	b, c := 20, 30
	fmt.Println("a= ", a, " b= ", b, " c= ", c)
	fmt.Printf("a = %d, b=%d ,c =%d \n", a, b, c)
	i, j := 10, 20
	i, j = j, i
	//匿名变量　_

	fmt.Printf("i= %d,j=%d\n", i, j)
	c, _, e := test()
	fmt.Printf("c=%d,e=%d\n", c, e)
}
