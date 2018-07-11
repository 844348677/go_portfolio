// 03_http
package main

import (
	"fmt"
	"net"
)

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("dial err = ", err)
		return
	}
	defer conn.Close()
	requestBuf := "GET /go HTTP/1.1\nHost: 127.0.0.1:8000\nUser-Agent: Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:45.0) Gecko/20100101 Firefox/45.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Language: en-US,en;q=0.5\nAccept-Encoding: gzip, deflate\nConnection: keep-alive\n\n"
	//requestBuf := "GET /goxxxx HTTP/1.1\nHost: 127.0.0.1:8000\nUser-Agent: Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:45.0) Gecko/20100101 Firefox/45.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Language: en-US,en;q=0.5\nAccept-Encoding: gzip, deflate\nConnection: keep-alive\n\n"

	//先发请求包　服务器才会回　响应包
	conn.Write([]byte(requestBuf))

	//接收服务器回复的响应包
	buf := make([]byte, 1024*4)
	n, err1 := conn.Read(buf)
	if n == 0 {
		fmt.Println("Read err = ", err1)
		return
	}

	//打印响应报文
	fmt.Printf("#%v#", string(buf[:n]))

	fmt.Println("Hello World!")
}
