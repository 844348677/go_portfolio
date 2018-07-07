package main

import (
	"fmt"
)

//无参数　无返回值　的函数定义
func MyFunc() {
	a := 666
	fmt.Println(" a = ", a)
}

//有参数　无返回值　函数　普通参数列表
func MyFunc01(a int) {
	fmt.Println("a = ", a)
}
func MyFunc02(a int, b int) {
	fmt.Println("a = %d , b = %d\n", a, b)
}
func MyFunc03(a, b int) {
	fmt.Println("a = %d , b = %d\n", a, b)
}

//　...int 这样的类型　...type 不定参数 传递的实参可以是０或多个
// 不定参数，只能设置　形参最后一个
func MyFunc04(args ...int) {
	// 获取用户传递参数的个数
	fmt.Println("len(args) = ", len(args))
	//　range返回两个值，下标　和对应的数
	for i, v := range args {
		fmt.Println(i, " ", v)
	}
}
func myfunc(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}
}
func myfunc2(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}
}
func test(args ...int) {
	//全部元素传递给ｍｙｆｕｎｃ
	myfunc(args...)
	// ！判断长度
	//　只把后两个参数传递给另一个函数
	myfunc2(args[len(args)-2:]...)
	myfunc2(args[:2]...)
}

//　无参数，有一个返回值 通过return中断函数　通过return 返回
func myfunc01() int {
	return 666
}

//　给返回值起一个变量名
func myfunc02() (result int) {
	result = 777
	return
}
func myfunc03() (int, int, int) {
	return 1, 2, 3
}
func myfunc04() (a, b, c int) {
	a, b, c = 4, 5, 6
	return
}
func main() {
	MyFunc()
	MyFunc01(222)
	//　调用函数　传递的参数　叫实参
	//　定义函数时　在函数后面（）中定义的参数　叫形参　实参传递给形参

	MyFunc04(1, 2, 3, 4)
	MyFunc04()
	test(5, 6, 7)
	//test()
	fmt.Println("a = ", myfunc01())
	c := myfunc02()
	fmt.Println(" c = ", c)

	e, d, f := myfunc04()
	fmt.Println("e= %d,d=%d,f=%d \n", e, d, f)
}
