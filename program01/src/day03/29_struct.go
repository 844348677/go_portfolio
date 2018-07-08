// 29_struct.go
package main

import (
	"fmt"
)

//定义一个结构体类型
type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	//顺序初始化　每个程序必须初始化
	var s1 Student = Student{1, "mike", 'm', 18, "bj"}
	fmt.Println("s1 = ", s1)
	//指定成员初始化，没有初始化的成员，自动赋值为０
	s2 := Student{name: "mike", addr: "bj"}
	fmt.Println("s2 = ", s2)

	var p1 *Student = &Student{1, "mike", 'm', 18, "bj"}
	fmt.Println("*p1 = ", *p1)

	p2 := &Student{name: "mike", addr: "bj"}
	fmt.Printf("p2 type is %T \n", p2)
	fmt.Println("p2 = ", p2)

	//定义一个结构体，普通变量
	var s Student

	//操作成员　使用.操作符
	s.id = 1
	s.name = "mike"
	s.sex = 'm'
	s.age = 18
	s.addr = "bj"
	fmt.Println("s = ", s)

	//指针有合法指向后　才操作成员
	//先定义一个普通结构体变量
	var a Student
	//再定义一个指针变量　保存ｓ的地址
	var p *Student
	p = &a
	//　通过指针操作成员　p.id 和　(*p).id 完全等价　只能使用.操作符
	p.id = 18
	(*p).name = "mike"
	p.sex = 'm'
	p.age = 18
	p.addr = "bj"
	fmt.Println("p = ", p)

	//通过new来申请一个结构体
	pn := new(Student)
	pn.id = 1
	pn.name = "mike"
	pn.sex = 'm'
	pn.age = 18
	pn.addr = "bj"
	fmt.Println(" pn = ", pn)

	a1 := Student{1, "mike", 'm', 18, "bj"}
	a2 := Student{1, "mike", 'm', 18, "bj"}
	a3 := Student{2, "mike", 'm', 18, "bj"}
	fmt.Println("s1 == s2 ", a1 == a2)
	fmt.Println("s1 == s3 ", a1 == a3)

	//同类行的两个结构体变量　可以相互赋值
	var tmp Student
	tmp = a3
	fmt.Println(" tmp = ", tmp)

	//函数传递　值传递　引用传递
	stu := Student{1, "mike", 'm', 18, "bj"}
	test01(stu)
	//值传递，形参无法改实参
	fmt.Println(" main: ", stu)

	test02(&stu) //地址　传递　引用　形参能改实参
	//操作指针所指向的内存
	fmt.Println("main: ", stu)
}

func test01(s Student) {
	s.id = 666
	fmt.Println(" test01: ", s)
}
func test02(p *Student) {
	p.id = 666
	fmt.Println(" test02: ", p)
}
