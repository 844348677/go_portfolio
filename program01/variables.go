// variables.go
package main

import (
	"fmt"
)

func main() {
	var a int
	//导入包　定义变量　要使用
	//　声明没使用变量　默认值为0
	//　同一个{}　声明变量是唯一的
	var b, c int
	fmt.Println("a= ", a)
	fmt.Println(b, c)
	a = 10 //变量的赋值

	//变量初始化, 声明变量的同时，同时赋值
	var d int = 10 //初始化　声明变量的时候赋值
	d = 20         //赋值
	fmt.Println(d)

	//short variable ...
	e := 30
	//%T 打印变量所属类型
	fmt.Printf("c type is %T\n", e)

	//赋值　赋值前必须先声明变量
	var n int
	n = 10
	fmt.Println(n)
	//先声明　ｂ的类型　再给ｂ赋值为２０
	m := 20
	//　不能在次　:=
	fmt.Println("m = ", m)
}
