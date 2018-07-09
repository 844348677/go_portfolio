// 07_strconv
package main

import (
	"fmt"
	"strconv"
)

func main() {
	//转换为字符串后追加到字节数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	//第二个数为要追加的数，第３个为指定１０进制方式追加
	slice = strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "abcgohello")
	fmt.Println("slice = ", string(slice))
	fmt.Println("len(slice) = ", len(slice))

	//其他类型转换为字符串
	var str string
	str = strconv.FormatBool(false)
	fmt.Println("str = ", str)
	//　'f'　指打印格式，以小数方式　－１指小数点位数（紧缩模式）　６４以ｆｌｏａｔ６４处理
	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	fmt.Println("str = ", str)
	str = strconv.Itoa(6666)
	fmt.Println("str = ", str)

	//字符串转其他类型
	var flag bool
	var err error
	flag, err = strconv.ParseBool("true")
	if err == nil {
		fmt.Println("flag = ", flag)
	} else {
		fmt.Println("error = ", err)
	}

	//把字符串转换为整型
	a, _ := strconv.Atoi("567")
	fmt.Println("a = ", a)
}
