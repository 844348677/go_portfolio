// 20_variables.go
package main

import (
	"fmt"
)

func test() {
	a := 20
	fmt.Println("a = ", a)
}
func test01() {
	fmt.Println("a = ", a)
}

//定义在函数外部的变量是全局变量
var a int

func main() {
	a = 10
	fmt.Println(" a = ", a)
	test()
	test01()
	fmt.Printf("%T\n", a)
	var a byte
	//局部变量
	//不同作用域，允许定义同名变量
	//　使用变量　就近原则
	fmt.Printf("%T\n", a)
	//定义在{}里面的变量就是局部变量
	//作用域,变量其作用的范围
	{
		i := 10
		fmt.Println("i = ", i)
	}
	if flag := 3; flag == 2 {
		fmt.Println(" flag = ", flag)
	} else {
		fmt.Println("flag = ", flag)
	}
	//局部变量离开作用域　自动释放
}
