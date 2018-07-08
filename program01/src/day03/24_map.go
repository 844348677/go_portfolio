// 24_map.go
package main

import (
	"fmt"
)

func test(m map[int]string) {

	delete(m, 1) //删除没有的键值也没有报错
	delete(m, 2)
}

func main() {
	var m1 map[int]string
	fmt.Println("m1 = ", m1)
	//对于map只有len,没有cap
	fmt.Println("len = ", len(m1))
	//可以通过ｍａｋｅ创建
	m2 := make(map[int]string)
	fmt.Println("m2 = ", m2)
	//可以通过ｍａｋｅ创建，可以指定长度
	m3 := make(map[int]string, 2)
	m3[1] = "mike"
	m3[2] = "go"
	m3[3] = "c++"
	fmt.Println("m3 = ", m3)
	fmt.Println("len(m3) = ", len(m3))
	//顺序　无序

	//　键值唯一　？？？
	//m4 := map[int]string{1: "aa", 1: "bb", 2: "cc"}
	//fmt.Println("m4 = ", m4)
	fmt.Println("Hello World!")

	map1 := map[int]string{1: "aa", 2: "bb"}
	fmt.Println("map1 = ", map1)
	map1[1] = "c++"
	fmt.Println("map1 = ", map1)
	map1[3] = "go"
	fmt.Println("map1 = ", map1)

	//删除 delete传变量　键值
	delete(map1, 1)
	//删除ｋｅｙ　为１的内容
	fmt.Println("map1 = ", map1)

	//ｍａｐ　做参数
	//　在函数内部删除某个ｋｅｙ
	test(map1)
	fmt.Println("m = ", map1)
}
