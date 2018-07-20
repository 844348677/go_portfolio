package main

import (
	"fmt"
)

func main() {
	fmt.Println("aaa")
	var slice1 []int
	slice2 := make([]int, 0)
	slice3 := []int{}
	fmt.Println(slice1 == nil)

	fmt.Println(slice2 == nil)
	fmt.Println(slice3 == nil)

	source := []string{"aa", "bb", "cc", "dd", "ee"}
	slice0 := source[2:3:3]
	fmt.Println(&slice0[0])
	fmt.Println(source)
	fmt.Println(slice0)
	// cap 容量扩增　之后　ｓｌｉｃｅ０　复制元素，换地址了
	slice0 = append(slice0, "nn")
	fmt.Println(&source[2])
	fmt.Println(&slice0[0])
	fmt.Println(slice0[0])
	slice0[0] = "mm"
	fmt.Println(source)
	fmt.Println(slice0)

}
