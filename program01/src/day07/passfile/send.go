package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func SendFile(path string, conn net.Conn) {
	//以只读方式打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open err = ", err)
		return
	}

	buf := make([]byte, 1024*4)

	//读文件内容　读多少就发多少　一点不差

	for {
		//从文件读取内容
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("f.Read err = ", err)
			}
			return
		}
		//发送内容
		conn.Write(buf[:n])

	}
}

func main() {
	//提示输入文件:
	fmt.Println("请输入需要传输的文件: ")
	var path string
	fmt.Scan(&path)

	//获取文件名 info.Name()
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//主动连接服务器
	conn, err1 := net.Dial("tcp", "127.0.0.1:8000")
	if err1 != nil {
		fmt.Println("net.Dial err1 = ", err1)
		return
	}
	defer conn.Close()

	//给接收放　先发送文件名 文件名不是路径
	_, err2 := conn.Write([]byte(info.Name()))
	if err2 != nil {
		fmt.Println("net.Dial err1 = ", err1)
		return
	}

	//接收对方的回复　如果回复”ok"　说明对方准备好　可以法文件
	var n int
	buf := make([]byte, 1024)
	n, err3 := conn.Read(buf)
	if err3 != nil {
		fmt.Println("net.Dial err1 = ", err3)
		return
	}
	if "ok" == string(buf[:n]) {
		//发送文件内容　路径
		SendFile(path, conn)
	}
}
