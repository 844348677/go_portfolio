// 14_func.go
package main

import (
	"fmt"
)

func test01() int {
	//函数被调用时，ｘ才分配空间，才初始化为０
	var x int //没有初始化　值为０
	x++
	//函数调用完毕，ｘ自动释放
	return x * x
}

//　函数的返回值是一个匿名函数　返回一个函数类型
func test02() func() int {
	var x int

	//独立的空间　ｘ的声明周期还存在
	return func() int {
		x++
		return x * x
	}
}

func main() {
	a := 10
	str := "make"
	func() {
		//闭包以引用方式　捕获外部变量　里面改了，外面也改了
		//不关心变量常量超出作用于
		a = 666
		str = "go"
		fmt.Printf("内部: a = %d,str = %s\n", a, str)
	}()
	//直接调用
	fmt.Printf("外部　a=%d , str = %s \n", a, str)

	fmt.Println(test01())
	fmt.Println(test01())
	fmt.Println(test01())
	//返回值为一个匿名函数，返回一个函数类型，通过ｆ来调用返回的匿名函数
	//闭包是穷人的类，类是穷人的闭包
	//只要闭包还在使用它，这些变量就还会存在
	f := test02()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	f2 := test02()
	fmt.Println(f2())
}
