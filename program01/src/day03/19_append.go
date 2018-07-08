// 19_append.go
package main

import (
	"fmt"
)

func main() {
	s1 := []int{}
	fmt.Printf("len = %d , cap =  %d\n", len(s1), cap(s1))

	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))
	fmt.Println("s1 = ", s1)

	s2 := []int{1, 2, 3}
	fmt.Println("s2 = ", s2)
	s2 = append(s2, 5)
	s2 = append(s2, 5)
	s2 = append(s2, 5)
	fmt.Println(" s2 = ", s2)
	//　以２倍容量扩容
	s := make([]int, 0, 1)
	oldCap := cap(s)
	for i := 0; i < 8; i++ {
		s = append(s, i)
		if newCap := cap(s); oldCap < newCap {
			fmt.Printf("cap: %d     %d\n", oldCap, newCap)
			oldCap = newCap
		}

	}
}
