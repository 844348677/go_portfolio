package main

//自动生成　ｐｋｇ　平台的静态库
// go install
//　ｂｉｎ　ｐｋｇ

import (
	"calc"
	"fmt"
)

//先执行ｉｎｉｔ　再执行　ｍａｉｎ函数
func init() {
	fmt.Println("this is main init")
}

// go env 查看go相关的环境路径
//　同一个目录，调用别的文件函数,直接调用即可，无需包名引用

func main() {
	Test()
	a := calc.Add(10, 20)
	fmt.Println(a)
	b := calc.Minus(5, 3)
	fmt.Println("b = ", b)
}
