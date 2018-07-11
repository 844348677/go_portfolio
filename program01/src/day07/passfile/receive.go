// receive
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

//接收文件内容
func RecvFile(fileName string, conn net.Conn) {
	//新建文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err = ", err)
		return
	}
	buf := make([]byte, 1024*4)

	//接收多少　，写多少　一点不差
	for {
		//接收对方发过来的文件内容
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read err = ", err)
			}
			return
		}
		if n == 0 {
			fmt.Println(" n == 0 文件接收完毕")
			return
		}
		f.Write(buf[:n])
	}
}

// ..1.png
func main() {
	//监听
	listenner, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("os.Stat err = ", err)
		return
	}
	defer listenner.Close()

	//阻塞等待用户连接
	conn, err1 := listenner.Accept()
	if err1 != nil {
		fmt.Println("listenner.Accept err1 = ", err1)
		return
	}

	defer conn.Close()

	buf := make([]byte, 1024)
	var n int
	n, err2 := conn.Read(buf) //读取对方发送的文件名
	if err2 != nil {
		fmt.Println("conn.Read err = ", err2)
		return
	}
	fileName := string(buf[:n])

	//回复　ok
	conn.Write([]byte("ok"))

	//接收文件内容
	RecvFile(fileName, conn)

	fmt.Println("Hello World!")
}
