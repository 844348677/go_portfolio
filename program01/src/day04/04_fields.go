// 04_fields.go
package main

import (
	"fmt"
)

type Person struct {
	name string
	sex  byte //字符类型
	age  int
}

//自定义类型，给类型改名
type mystr string
type Student struct {
	//结构体匿名字段
	Person //只有类型　没有名字　匿名字段　继承了Ｐｅｒｓｏｎ的成员
	//基础类型的匿名字段
	int
	mystr
}

func main() {
	s := Student{Person{"mike", 'm', 18}, 666, "aaa"}
	fmt.Printf("s = %+v\n", s)
	s.Person = Person{"go", 'm', 22}
	fmt.Println(s.name, s.age, s.sex, s.int, s.mystr, s.Person)
}
