// 24_switch.go
package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("input")
	fmt.Scan(&num)

	switch num {
	case 1:
		fmt.Printf("1 %d \n", num)
		//break       //可写可不写 break关键字，跳出
		fallthrough //不跳出switch语句，后面无条件执行 ???????
	case 2:
		fmt.Printf("2 %d \n", num)
	case 3, 4, 5:
		fmt.Printf("3,4,5 %d \n", num)
	default:
		fmt.Println("xxxx")

	}
	fmt.Println("Hello World!")

	switch num := 1; num {
	case 1:
		fmt.Println("1")
	}

	score := 85
	switch {
	case score > 90:
		fmt.Println("优秀")
	case score > 80:
		fmt.Println("良好")
	case score > 70:
		fmt.Println("差")
	default:
		fmt.Println("其他")
	}
}
