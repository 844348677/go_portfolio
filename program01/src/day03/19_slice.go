// 19_slice.go
package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//新切片
	s1 := a[2:5]
	fmt.Println("s1 = ", s1)

	s1[1] = 666
	fmt.Println("s1 = ", s1)
	fmt.Println("a = ", a)

	//另外的新切片
	// 0 1 2 3 4 5 6 7 8 9
	//       2 3 4
	//             4 5 6 7 8
	s2 := s1[2:7]
	fmt.Println("s2 = ", s2)
	s2[2] = 777
	fmt.Println("s2 = ", s2)
	fmt.Println("a = ", a)
	fmt.Println("s1 = ", s1)

}
