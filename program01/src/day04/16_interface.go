// 16_interface.go
package main

import (
	"fmt"
)

type Humaner interface {
	sayhi()
}
type Personer interface {
	Humaner
	sing(lrc string)
}

type Student struct {
	name string
	id   int
}

func (tmp *Student) sayhi() {
	fmt.Printf("Student [%s , %d] sayhi \n", tmp.name, tmp.id)
}
func (tmp *Student) sing(lrc string) {
	fmt.Println("Student 在唱着: ", lrc)
}
func main() {
	// 定义一个接口类型变量
	var i Personer
	s := &Student{"mike", 666}
	// 继承过来的方法
	i = s
	i.sayhi()
	i.sing(" 歌歌")
	fmt.Println("Hello World!")

	//接口的转化
	var interface_personer Personer
	interface_personer = &Student{"mike", 666}
	var interface_humaner Humaner

	//interface_personer = interface_humaner
	//error
	//
	interface_humaner = interface_personer
	interface_humaner.sayhi()
}
