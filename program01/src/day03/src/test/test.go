package test

import "fmt"

//如果首字母是小写，只能在同一个包使用
func myFunc() {
	fmt.Println("this is myFunc()")
}

func MyFunc() {
	fmt.Println("this is MyFunc()")
}

type stu struct {
	id int
}

type Stu struct {
	id int
	Id int
}
