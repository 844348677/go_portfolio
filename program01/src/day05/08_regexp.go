// 08_regexp
package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "abc azc a7c aac 888 a9c tac"
	//1. 解释规则, 它会解析正则表达式　如果成功 返回解析器
	//reg1 := regexp.MustCompile(`a.c`)
	//reg1 := regexp.MustCompile(`a[0-9]c`)
	reg1 := regexp.MustCompile(`a\dc`)
	if reg1 == nil { //解析失败　返回nil
		fmt.Println("err = ", reg1)
		return
	}
	// 2. 根据规则提取　关键信息 -1代表解析所有
	result1 := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result1 = ", result1)
	result2 := reg1.FindAllStringSubmatch(buf, 1)
	fmt.Println("result1 = ", result2)

	buf2 := "43.14 567 agsdg 1.23 7. 8.99 1sdljgl 6.66 7.8"
	//解释正则表达式　＋匹配前一个字符的１次或多次
	reg := regexp.MustCompile(`\d+\.\d+`)
	if reg == nil {
		fmt.Println("MustCompile err")
		return
	}
	//提取关键信息
	//result := reg.FindAllString(buf2, -1)
	result := reg.FindAllStringSubmatch(buf2, -1)
	fmt.Println("result = ", result)

}
