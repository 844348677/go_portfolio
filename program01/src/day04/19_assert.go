// 19_assert
package main

import (
	"fmt"
)

type Student struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1           //int
	i[1] = "hello go " //string
	i[2] = Student{"mike", 666}

	//类型查询　类型断言
	// 第一个返回下标　第二个返回下标对应的值，data 分别是i[0],i[1],i[2]

	for index, data := range i {
		//第一个返回的是值　接口变量本身　第二个判断结果的真假
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d] 类型为int, 内容为　%d\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d] 类型是string, 内容为 %s\n", index, value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Printf("x[%d] 类型为Student, 内容为name = %s , id = %d\n", index, value.name, value.id)
		}
	}
	fmt.Println("Hello World!")
}
