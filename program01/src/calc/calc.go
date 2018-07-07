package calc

import (
	"fmt"
)

//不同目录　包名不一样
//　调用不同包里的函数　包名.函数()
//调用别的包的函数, 这个包函数的名字是小写　无法让别人调用，必须首字母大写
//首字母大写
func add(a, b int) int {
	return a + b
}
func Add(a, b int) int {
	return a + b
}
func Minus(a, b int) int {
	return a - b
}

//导入包　先执行包里的　init函数
func init() {
	fmt.Println("this is calc init")
}
