// 17_input.go
package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Hello World!")
	//　阻塞等待用户输入
	//fmt.Scanf("%d", &a)
	fmt.Scan(&a)
	fmt.Println("a = ", a)

	var flag bool
	flag = true
	fmt.Printf("flag = %d\n", flag)
	fmt.Printf("flag = %t\n", flag)
	//fmt.Printf("flag = %d\n", int(flag))
	//bool 类型不能转化为整型
	// 0就是假, 非０就是真
	//整型不能转化为ｂｏｏｌ类型
	// flag = bool(int)
	//这种不能转化的类型,叫不兼容类型
	var ch byte
	ch = 'a'
	var t int
	t = int(ch)
	//类型转换，把ch的值取出来,转成int再赋值
	fmt.Println("t= ", t)

	type bigint int64
	var x bigint //等价于　int64 a
	// main.bigint
	type (
		longint int64
		char    byte
	)
	var y longint = 11
	var z char = 'a'
	fmt.Printf("a type is %T\n", x)
	fmt.Printf("y = %d, z = %c\n", y, z)
	fmt.Println("4 > 3 ", 4 > 3)
	// !true !false
	//&& 都为真
	// ||　一个或两个为真，真
	//go语言中bool类型和int不兼容
	a = 11
	fmt.Println("0<= a && a<=10", 0 <= a && a <= 10)

	s := "111"
	if s == "111" {
		fmt.Println("2222")
	}
	//if 支持一个初始化语句，初始化语句和判断以分号分割
	if a := 101; a == 101 {
		fmt.Println("a==101", a)
	}
	if a == 10 {
		fmt.Println("a == 10")
	} else {
		fmt.Println("a != 10")
	}
	if a == 10 {
		fmt.Println("a == 10")
	} else if a > 10 {
		fmt.Println("a >10")
	} else if a < 10 {
		fmt.Println("a < 10 ")
	} else {
		fmt.Println("nononon")
	}

}
