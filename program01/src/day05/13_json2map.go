// 13_json2map
package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	//字段改成小写
	//相当于改名字　二次编码
	Company  string   `json:"company"` //此字段不会输出到屏幕
	Subjects []string `json:"subjects"`
	IsOk     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

func main() {
	jsonBuf := `
	{
	 "company": "it",
	 "subjects": [
	  "go",
	  "C++",
	  "Python",
	  "Test"
	 ],
	 "isok": true,
	 "price": 666.66
	}
	`
	var tmp IT
	//第二个参数　去地址　否则无法改值
	err := json.Unmarshal([]byte(jsonBuf), &tmp)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("tmp = ", tmp)
	fmt.Printf("tmp = %+v\n", tmp)

	type IT2 struct {
		Subjects []string `json:"subjects"`
	}
	var tmp2 IT2
	//只会解析所需要的字段，放到对应的位置
	err = json.Unmarshal([]byte(jsonBuf), &tmp2)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("tmp2 = %+v\n", tmp2)

	//创建一个ｍａｐ
	m := make(map[string]interface{}, 4)
	err = json.Unmarshal([]byte(jsonBuf), &m)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("m = %+v\n", m)

	var str string
	//str = m["company"]
	//str = string(m["company"]) //err 无法转换
	fmt.Println("str = ", str)

	//类型断言
	for key, value := range m {
		//fmt.Printf("%v ========> %v \n", key, value)
		switch data := value.(type) {
		case string:
			str = data
			fmt.Printf("map[%s] 的值类型为string ，value = %s\n ", key, str)
		case bool:
			fmt.Printf("map[%s] 的值类型为bool ，value = %v\n ", key, data)
		case float64:
			fmt.Printf("map[%s] 的值类型为float64 ，value = %f\n ", key, data)
		case []string:
			fmt.Printf("map[%s]的值类型为[]string ,value = %v\n", key, data)
		//空接口的slice 再解析反推
		case []interface{}:
			fmt.Printf("map[%s]的值类型为[]interface ,value = %v \n", key, data)
		}
	}
}
