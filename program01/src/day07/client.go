// client
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dail err = ", err)
		return
	}
	//main 调用完毕　关闭连接
	defer conn.Close()

	go func() {
		//从键盘输入内容　给服务器发送内容
		str := make([]byte, 1024)
		for {

			//从键盘读取内容　放在str
			n, err := os.Stdin.Read(str)

			if err != nil {
				fmt.Println("os.Stdin.Read err = ", err)
				return
			}
			//把输入的内容发给服务器发送
			conn.Write(str[:n])
		}
	}()
	//接收　服务器回复的数据
	//slice 缓存
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		//接收服务器的请求
		if err != nil {
			fmt.Println("conn.Read err = ", err)
			return
		}
		//转换成字符串　再打印
		fmt.Println(string(buf[:n])) //打印接收到的内容
	}

	fmt.Println("Hello World!")
}
