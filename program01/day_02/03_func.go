// 03_func.go
package main

import (
	"fmt"
)

func test(a int) {
	//函数终止调用的条件
	if a == 1 {
		fmt.Println("aａ = ", a)
		return
	}
	//函数调用自身
	test(a - 1)
	//　先进后出
	fmt.Println("a = ", a)
}
func test01() (sum int) {
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return

}
func test02(i int) int {
	if i == 1 {
		return 1
	}

	return i + test02(i-1)
}
func test03(i int) int {
	if i == 100 {
		return 100
	}
	return i + test03(i+1)
}
func add(a, b int) int {
	return a + b
}
func minus(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}

//　回调函数，函数有一个参数，是函数类型
func Calc(a, b int, fTest FuncType) (result int) {
	fmt.Println("Calc")
	//　fTest还没实现
	result = fTest(a, b)
	return
}

// FuncType 　是一个函数类型 　函数的　参数和返回值
type FuncType func(int, int) int

func main() {
	test(3)
	fmt.Println("main")
	var sum int
	sum = test02(100)
	fmt.Println("sum = ", sum)

	//传统调用方法
	var result int
	result = add(1, 1)
	//函数也是一种类型，通过ｔｙｐｅ给一个函数类型起名
	var fTest FuncType
	fTest = add
	result = fTest(10, 20)
	fmt.Println("result = ", result)

	fTest = minus
	result = fTest(10, 5)

	fmt.Println("result = ", result)
	//　多态　多种形态，调用同一个接口，不同的表现，　实现不同的表现
	//现在有想法　后面再实现功能
	a := Calc(1, 1, add)
	a = Calc(1, 1, mul)
	fmt.Println("a = ", a)
}
