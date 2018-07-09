package main

import (
	"fmt"
)

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[0:2]
	fmt.Printf(" %T \n", slice)
	fmt.Println(slice)
	array[0] = 666
	fmt.Println(slice)
}
