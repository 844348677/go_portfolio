// 12_mapjson
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//创建一个map
	m := make(map[string]interface{}, 4)
	m["company"] = "itcast"
	m["subjects"] = []string{"Go", "C++", "Python", "Test"}
	m["isok"] = true
	m["price"] = 666.66

	result, err := json.Marshal(m)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("result = ", string(result))
	fmt.Println("Hello World!")
}
