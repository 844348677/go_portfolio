// 01_struct.go
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
}

func main() {
	var s1 Student = Student{Person{"mike", 'm', 18}, 1, "bj"}
	fmt.Println("s1 = ", s1)

	s2 := Student{Person{"mike", 'm', 18}, 1, "bj"}
	fmt.Println(" s2 = ", s2)
	// %+v 显示更详细
	fmt.Printf("s2 = %+v\n", s2)

	//指定成员初始化　没有初始化的成员自动赋值为０
	s3 := Student{id: 1}
	fmt.Printf("s3 = %+v\n", s3)

	s4 := Student{Person: Person{name: "mike"}, id: 1}
	fmt.Printf("s4 = %+v\n", s4)

	//s5 := Student{"mike", 'm', 18, 1, "bj"}

	s1.name = "bb"
	s1.sex = 'f'
	s1.age = 22
	s1.id = 666
	s1.addr = "sz"

	s1.Person = Person{"go", 'm', 18}
	fmt.Println(s1.name, s1.sex, s1.addr, s1.id, s1.addr)

}
