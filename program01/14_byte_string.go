// 14_byte_string.go
package main

import (
	"fmt"
)

func main() {
	var ch byte
	var str string

	//字符 一个字符，转义字符
	ch = 'a'
	//字符串　一个或多个字符组成
	//　字符串　隐藏了一个结束字符　'\0' 　ｃ++ 　好像也有这个设置
	str = "a"
	fmt.Println("ch = ", ch)
	fmt.Println("str = ", str)
	str = "hello go"
	//　操作字符串的某个字符　从０开始操作
	fmt.Printf("str[0] = %c,str[1] = %c\n", str[0], str[1])

	var t complex128
	t = 2.1 + 3.14i
	fmt.Println("t= ", t)
	t2 := 3.3 + 4.4i
	fmt.Println("t2= ", t2)

}
