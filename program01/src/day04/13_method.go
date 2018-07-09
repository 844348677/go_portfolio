// 13_method.go
package main

import (
	"fmt"
)

type Person struct {
	name string
	sex  byte //字符类型
	age  int
}

func (p Person) SetInfoValue() {
	fmt.Printf("setinfovalue : %p , %v\n", &p, p)
}

//　接收者　为指针变量　引用传递
func (p *Person) SetInfoPointer() {
	fmt.Printf("SetInfoPointer : %p , %v\n", p, *p)
}

func main() {
	p := Person{"mike", 'm', 18}
	fmt.Printf("main : %p , %v\n", &p, p)
	//引用传递
	p.SetInfoPointer()

	//　保存　方法入口地址
	pFunc := p.SetInfoPointer //这个就是方法值　调用函数时　无需再传递接收者，隐藏了接收者
	pFunc()

	// 值传递　不同的地址
	vFunc := p.SetInfoValue
	vFunc()

	//方法值　f := p.SetInfoPointer
	//方法表达式
	f := (*Person).SetInfoPointer
	//需要　显示把接收者传递过去 p.SetInfoPointer()
	f(&p)

	f2 := (Person).SetInfoValue
	f2(p)

}
