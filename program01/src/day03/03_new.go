// 03_new.go
package main

import (
	"fmt"
)

func main() {
	//a := 10

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
}
