package main

import "fmt"
import "time"

func newTask() {
	for {
		fmt.Println("this is a newTash")
		time.Sleep(time.Second)
	}
}

func main() {
	go newTask() //新建一个goroutine 新建一个任务
	//　只要go 开新的　goroutine 执行任务

	go func() {
		i := 0
		for {
			i++
			fmt.Println("　ｉ＝　", i)
			time.Sleep(time.Second)
		}
	}()
	//　main 退出了　其他go　也退出
	i := 0
	for {
		i++
		fmt.Println("this is a main goroutine ")
		//延时一秒
		time.Sleep(time.Second)
		if i == 4 {
			break
		}
	}

}
