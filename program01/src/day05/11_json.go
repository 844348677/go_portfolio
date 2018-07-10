// 11_json
package main

import (
	"encoding/json"
	"fmt"
)

//成员变量名　首字母必须大写
type IT struct {
	//字段改成小写
	//相当于改名字　二次编码
	Company  string `json:"company"` //此字段不会输出到屏幕
	Subjects []string
	IsOk     bool `json:",string"`
	Price    float64
}

func main() {
	//定义一个结构体变量，同时初始化
	s := IT{"it", []string{"go", "C++", "Python", "Test"}, true, 666.66}

	//编码　根据内容生成json文本
	buf, err := json.Marshal(s)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	// {"Company":"it","Subjects":["go","C++","Python","Test"],"IsOk":true,"Price":666.66}
	fmt.Println("buf = ", string(buf))

	//　格式化　编码
	buf2, err2 := json.MarshalIndent(s, "", " ")
	if err2 != nil {
		fmt.Println("err = ", err2)
		return
	}
	fmt.Println("buf = ", string(buf2))
}
