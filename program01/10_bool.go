// 10_bool.go
package main

import (
	"fmt"
)

func main() {
	//声明变量
	var a bool
	fmt.Println(a)
	a = true

	//
	b := false
	fmt.Println(b)
	var c = false
	fmt.Println(c)

	var f1 float32
	f1 = 3.14
	fmt.Println(f1)

	f2 := 3.14
	// f2 type is float64
	fmt.Printf("f2 type is %T\n", f2)

	// ascii
	// 97 'a' 65'A' byte 一个字节　ａｓｃｉｉ
	var ch byte
	ch = 97
	fmt.Println("ch = ", ch)
	//格式化输出　%c以字符方式打印　％ｄ以整数方式
	fmt.Printf("%c, %d\n", ch, ch)

	ch = 'a' //字符是单引号的
	fmt.Printf("%c, %d\n", ch, ch)
	//　大小写转换 之间相差32 小写大
	fmt.Printf("大写: %d, 小写: %d\n", 'A', 'a')
	fmt.Printf("%c\n", 'A'+32)
	fmt.Printf("%c\n", 'a'-32)

	//　反斜杠开头的字符是转义字符
	fmt.Printf("1%c", '\n')
	fmt.Printf("2")

}
