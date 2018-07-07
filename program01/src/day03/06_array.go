// 06_array.go
package main

import (
	"fmt"
)

func main() {
	//数组，同一类型的集合
	var id [50]int
	for i := 0; i < len(id); i++ {
		id[i] = i + 1
		fmt.Printf("id[%d] = %d \n", i, id[i])
	}
	fmt.Println("Hello World!")

	// [10]int和[5]int不同类型
	//数组元素个数
	var a [10]int
	var b [5]int
	fmt.Printf("len(a) = %d , len(b) = %d \n", len(a), len(b))
	//n := 10
	// error
	//var c [n]int
	//　操作数组元素　０到len(n)-1　下标
	a[0] = 1
	i := 1
	a[i] = 2
	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}
	for i, data := range a {
		fmt.Printf("a[%d] = %d\n", i, data)
	}
	//数组初始化
	//声明定义　同时赋值　初始化
	var a1 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(" a1 =  ", a1)

	a2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(" a2 =  ", a2)
	//　部分初始化　没有初始化的元素　自动赋值为０
	a3 := [5]int{1, 2, 3}
	fmt.Println("a3 = ", a3)
	//　指定某个元素初始化
	d := [5]int{2: 10, 4: 20}
	fmt.Println("d= ", d)

	//有多少[] 多少维，有多少[]　多少循环
	var f [3][4]int
	k := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			k++
			f[i][j] = k
			fmt.Printf("f[%d][%d] = %d ", i, j, f[i][j])
		}
		fmt.Println()
	}
	fmt.Println("f = ", f)
	//有三个元素，每个元素又是一维数组
	f2 := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println("f2 = ", f2)
	//　部分初始化，不初始化的值为０
}
