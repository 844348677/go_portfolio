// 03_new.go
package main

import (
	"fmt"
)

func main() {
	var p *int
	//指向一个合法内存
	// p是*int 指向　ｉｎｔ类型
	p = new(int)
	*p = 666
	//动态分配空间　ｇｏ不需要释放
	fmt.Println("*p = ", *p)
	q := new(int)
	*q = 777
	fmt.Println("*q = ", *q)

	a, b := 10, 20
	//变量本身传递　值传递
	swap(a, b)
	fmt.Printf("swap: a= %d, b = %d \n", a, b)
	swap2(&a, &b)
	fmt.Printf("swap: a= %d, b = %d \n", a, b)

}
func swap(a, b int) {
	a, b = b, a
	fmt.Printf("swap: a= %d, b = %d \n", a, b)
}
func swap2(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
}
