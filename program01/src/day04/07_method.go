// 07_method.go
package main

import (
	"fmt"
)

type Person struct {
	name string
	sex  byte //字符类型
	age  int
}

//带有接收者的函数　叫方法
func (tmp Person) PrintInfo() {
	fmt.Println(" tmp = ", tmp)
}

//通过一个函数，给成员赋值
func (p *Person) SetInfo(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
}

type long int

func (tmp long) test() {

}

// receiver type 不能是指针或接口
type pointer *int

//不支持重载
type char byte

//接收者不同类型　函数名字同名　也是不同的方法　不是重复定义函数
func (tmp char) test() {

}
func main() {
	//定义同时初始化
	p := Person{"mike", 'm', 18}
	p.PrintInfo()

	//定义一个结构体变量
	var p2 Person
	(&p2).SetInfo("yoyo", 'f', 22)
	p2.PrintInfo()

	var s1 Person
	s1.name = "go"
	s1.age = 18
	s1.sex = 'm'
	fmt.Printf(" main: &p = %p \n", &s1)

	// 值传递
	s1.SetInfoValue("mike", 'm', 18)
	fmt.Println("s1 = ", s1)

	//引用
	(&s1).SetInfoPointer("mike", 'm', 18)
	fmt.Println(" s1 = ", s1)
}

//接收者　为普通变量　非指针　值传递　一份ｃｏｐｙ
func (p Person) SetInfoValue(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Printf(" setinfovalue &p = %p \n", &p)
	fmt.Println(" p = ", p)
}

//　接收者　为指针变量　引用传递
func (p *Person) SetInfoPointer(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Printf(" setinfopointer &p = %p \n ", p)
}
