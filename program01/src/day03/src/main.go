package main

import (
	"fmt"
	"test"
)

func main() {
	//　包名.函数名
	test.MyFunc()

	//　包名.结构体名
	var s test.Stu
	fmt.Println(" test.Stu = ", s)
	//s.id = 666
	s.Id = 666
	fmt.Println(" test.Stu = ", s)
}
