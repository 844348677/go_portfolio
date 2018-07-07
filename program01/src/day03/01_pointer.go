package main

import (
	"fmt"
)

func main() {
	var a int = 10
	//　每个变量有２层含义，变量的内存　变量的地址
	fmt.Printf(" a = %d \n", a) //变量的内存
	//变量的地址
	fmt.Printf("&a = %v \n", &a)

	//保存变量的地址　指针类型　*int保存int的地址　**int保存*int地址

	var p *int
	p = &a
	fmt.Printf("p = %v, &a = %v \n", p, &a)
	*p = 666
	// *p 操作的不是ｐ的内存，是ｐ所指向的内存　就是ａ
	fmt.Printf("*p = %v , a = %v \n", *p, a)

	var pointer *int
	p = nil
	fmt.Println("pointer = ", pointer)

	//不要操作没有合法指向地址　的指针
	// *p = 666

}
