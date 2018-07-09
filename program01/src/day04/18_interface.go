// 18_interface.go
package main

import (
	"fmt"
)

// 可以接收任意类型　任意个数的参数
func xxx(arg ...interface{}) {

}

func main() {
	// 空接口　万能类型　保存任意类型的值
	var i interface{} = 1
	fmt.Println("i = ", i)
	i = "abc"
	fmt.Println("i = ", i)
	// 空接口里面没有方法，所有类型都实现空接口
	fmt.Println("Hello World!")
}
