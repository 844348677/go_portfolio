// 15_slice.go
package main

import (
	"fmt"
)

func main() {
	//　初始化　，同时定义赋值
	s1 := []int{1, 2, 3, 4}
	fmt.Println("s1 = ", s1)
	//借助　ｍａｋｅ函数 make(slice type, len, cap)
	s2 := make([]int, 5, 10)
	fmt.Printf("len = %d, cap = %d \n", len(s2), cap(s2))
	//　没指定，长度和容量一致
	s3 := make([]int, 5)
	fmt.Printf("len = %d, cap = %d\n", len(s3), cap(s3))

	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// [low : high : max]
	//取下标　从low开始的元素, len = high-low , cap = max-low
	s4 := array[:] // [0:len(array):len(array)] 不指定容量　容量和长度一样
	fmt.Printf("len = %d, cap = %d\n", len(s4), cap(s4))
	// 操作某个元素，　和数组操作方式一样
	data := array[1]
	fmt.Println("data = ", data)

	s5 := array[3:6:7] // a[3],a[4],a[5] len = 6-3 = 3 cap = 7-3 =4
	fmt.Println("s5 = ", s5)
	fmt.Printf("len = %d, cap = %d\n", len(s5), cap(s5))

	s6 := array[:6] //从０开始　取六个元素　容量也是６
	fmt.Println("s6 = ", s6)
	fmt.Printf("len = %d , cap = %d\n", len(s6), cap(s6))

	s7 := array[3:] // 下标为３开始　到结尾
	fmt.Println("s7 = ", 7)
	fmt.Printf("len = %d , cap = %d\n", len(s7), cap(s7))
}

func main01() {
	a := []int{1, 2, 3, 0, 0}
	s := a[0:3:5]
	fmt.Println("s = ", s)
	//长度
	fmt.Println("len(s) = ", len(s))
	fmt.Println("cap(s) = ", cap(s))
	//　ｓｌｉｃｅ　０－３左闭右开　５－０为容量　ｃａｐ

	s2 := a[1:4:5]
	fmt.Println("s2 = ", s2)
	fmt.Println("len(s2) = ", len(s2))
	fmt.Println("cap(s2) = ", cap(s2))

	//切片和数组的区别
	//数组[]　里面的长度是固定的一个常量,数组不能修改长度　len和cap永远都是５
	b := [5]int{}
	fmt.Printf(" len = %d , cap = %d \n", len(b), cap(b))

	//切片 []里面为空　或者　...　切片的长度或者容量　不固定
	s3 := []int{}
	fmt.Printf(" len = %d , cap = %d \n", len(s3), cap(s3))
	//给切片末尾追加一个成员
	s3 = append(s3, 11)
	fmt.Printf(" len = %d , cap = %d \n", len(s3), cap(s3))
}
