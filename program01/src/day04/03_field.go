// 03_field.go
package main

import (
	"fmt"
)

type Person struct {
	name string
	sex  byte //字符类型
	age  int
}
type Student struct {
	Person //只有类型　没有名字　匿名字段　继承了Ｐｅｒｓｏｎ的成员
	id     int
	addr   string
	name   string //和ｐｅｒｓｏｎ同名了
}

func main() {
	//声明　定义一个变量
	var s Student
	//操作的是Ｓｔｕｄｅｎｔ还是ｐｅｒｓｏｎ中的ｎａｍｅ
	s.name = "mike"
	//　默认规则，　就近原则　如果能在本作用域找到次成员　操作此成员
	//　如果没有找到　找到继承的字段
	s.sex = 'm'
	s.age = 18
	s.addr = "bj"
	//　显示调用
	s.Person.name = "aa"
	fmt.Printf("s = %+v \n", s)
}
