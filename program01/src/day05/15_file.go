// 15_file
package main

import (
	"fmt"
	"os"
)

func main() {

	//os.Stdout.Close() //关闭后无法直接输出

	fmt.Println("Hello World!") //往标准输出设备　屏幕　写内容
	//标准设备文件　默认已经打开，用户可以直接使用

	os.Stdout.WriteString(" are u ok \n")

	os.Stdin.Close()
	//关闭后无法输入
	var a int
	fmt.Println("请输入　a: ")
	//从标准设备中　读取内容　放在ａ中
	fmt.Scan(&a)
	fmt.Println(" a = ", a)

}
