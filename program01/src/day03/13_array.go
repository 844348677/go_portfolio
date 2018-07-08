// 13_array.go
package main

import (
	"fmt"
)

//数组做函数参数　a值传递
//实参　数组每个元素　给形参数组　拷贝一份
//形参的数组　是实参的数组的复制
func modify(a [5]int) {
	a[0] = 666
	fmt.Println("modify a = ", a)
}

//　ｐ指向现实数组ａ　它是指向是数组　　数组指针
// *p　代表指针指向的内存　实参ａ
func modify2(p *[5]int) {
	(*p)[0] = 666
	fmt.Println("modify *a = ", *p)
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	modify(a)
	fmt.Println("main: a= ", a)

	modify2(&a)
	fmt.Println("main: a= ", a)
}
