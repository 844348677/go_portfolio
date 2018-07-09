// 06_string.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// "hellogo" 中是否包含　"hello" 包含返回ｔｒｕｅ　不包含返回ｆａｌｓｅ
	fmt.Println(strings.Contains("hellogo", "hello"))
	fmt.Println(strings.Contains("hellogo", "abc"))

	//Joins 组合
	s := []string{"abc", "hello", "mike", "go"}
	buf := strings.Join(s, "@")
	fmt.Println("buf = ", buf)

	//index 查找子串的位置
	fmt.Println(strings.Index("abcdhello", "hello"))
	//不存在返回　-1
	fmt.Println(strings.Index("abcdhello", "go"))

	buf = strings.Repeat("go", 3)
	fmt.Println("buf = ", buf)

	//split 以指定的分隔符拆分
	buf = "hello@abc@go@mike"
	s2 := strings.Split(buf, "@")
	fmt.Println("s2 = ", s2)
	fmt.Println("len(s2) = ", len(s2))

	buf = strings.Trim("       are u ok?          ", " ") //去掉两头空格
	fmt.Printf("buf = #%s#\n", buf)

	// 去掉空格　把元素放到切片中
	s3 := strings.Fields("       are u ok?          ")
	//fmt.Println("s3 = ", s3)
	for i, data := range s3 {
		fmt.Println(i, ",", data)
	}

}
