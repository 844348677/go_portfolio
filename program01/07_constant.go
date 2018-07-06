// 07_constant.go
package main

import (
	"fmt"
)

func main() {
	//变量　程序运行期间　可以改变
	//常量　程序运行期间，不可以改变　const
	const a int = 10
	//a = 20
	//常量值能不能修改
	const b = 10.0
	// 常量不能使用　:=
	fmt.Printf("b type is %T\n", b)
	fmt.Println(b)
	fmt.Println("a = ", a)

	var (
		n int
		m float64
	)
	n, m = 10, 3.14
	fmt.Println("n = ", n)
	fmt.Println("m= ", m)
	const (
		i int     = 10
		j float64 = 3.14
	)
	fmt.Println("i = ", i)
	fmt.Println("j= ", j)
}
