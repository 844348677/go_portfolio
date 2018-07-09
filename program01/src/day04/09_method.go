// 09_method.go
package main

import (
	"fmt"
)

type Person struct {
	name string
	sex  byte //字符类型
	age  int
}

//接收者　为普通变量　非指针　值传递　一份ｃｏｐｙ
func (p Person) SetInfoValue() {

	fmt.Println("setinfovalue")
}

//　接收者　为指针变量　引用传递
func (p *Person) SetInfoPointer() {
	fmt.Println("setinfopointer ")
}
func main() {
	//结构体的变量是一个指针变量　它能够调用哪些方法，这些方法就是一个集合，简称方法集
	p := &Person{"mike", 'm', 18}

	p.SetInfoPointer()
	(*p).SetInfoPointer()
	// 内部做的转换，先把指针ｐ　转成*p　后再调用
	p.SetInfoValue()
	(*p).SetInfoValue()

	//指针和非指针　ｒｅｃｅｉｖｅｒ　自动转化

	p2 := Person{"mike", 'm', 18}
	//内部　先把p 转化为&p　再调用  (&p).SetInfoPointer
	p2.SetInfoPointer()
	p2.SetInfoValue()

	s := Student{Person{"mike", 'm', 18}, 666, "bj"}
	//　ｐｅｒｓｏｎ的方法
	s.PrintInfo()
	//　调用就近原则　先找本作用域的方法，找不到再用继承的
	//显示调用　继承的方法
	s.Person.PrintInfo()
}

//　学生　结构体　继承ｐｅｒｓｏｎ字段，成员和方法都继承了
type Student struct {
	//匿名字段
	Person
	id   int
	addr string
}

// student 也实现了一个方法　这个方法和Ｐｅｒｓｏｎ方法同名　这种方法叫重写
func (tmp *Student) PrintInfo() {
	//tmp.Person.PrintInfo()
	fmt.Println("studnet : tmp = ", tmp)
}

func (tmp *Person) PrintInfo() {
	fmt.Printf(" name = %s, sex = %c , age = %d\n", tmp.name, tmp.sex, tmp.age)
}
