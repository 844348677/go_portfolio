// 02_func.go
package main

import (
	"fmt"
)

func MaxAndMin(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	return
}
func funcb() (b int) {
	b = 222
	fmt.Println("funcb b = ", b)
	return
}
func funca() (a int) {
	a = 111
	b := funcb()
	fmt.Println("funca b = ", b)
	fmt.Println("funca a = ", a)
	return
}
func main() {
	max, min := MaxAndMin(10, 20)
	fmt.Printf("max = %d, min = %d\n", max, min)
	fmt.Println("main func")
	a := funca()
	fmt.Println("main  a= ", a)
}
