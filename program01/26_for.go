// 26_for.go
package main

import (
	"fmt"
	"time"
)

func main() {
	// for 初始条件; 判断条件; 条件变化{
	//}
	sum := 0
	// i初始赋值为１　1<=100为真继续循环
	for i := 1; i <= 100; i++ {
		sum = sum + i
	}
	fmt.Println("sum = ", sum)

	// range  foreach
	str := "abc"
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}
	//range 返回两个值　元素的位置　和　元素的本身
	for i, data := range str {
		fmt.Printf("st[%d] = %c\n", i, data)
	}
	for i, _ := range str {
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}
	// break for switch select continue for
	//　for ,死循环
	i := 0
	for {
		i++
		time.Sleep(time.Second)
		if i == 5 {
			break
			//跳出for 循环，多嵌套循环，跳出最近的
			continue
			//跳过本次循环，继续执行循环
		}
		fmt.Println("i= ", i)
	}
	//goto 可以用在任何地方，但是不能跨函数使用
	fmt.Println("11111111")
	goto End
	fmt.Println("2222222")
End:
	fmt.Println("3333333")
}
