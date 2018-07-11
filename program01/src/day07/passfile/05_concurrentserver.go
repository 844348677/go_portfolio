// 05_concurrentserver
package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

//nc 127.0.0.1 8000 连接到服务器
// who 查询在线client
//

type Client struct {
	C    chan string //用于发送数据的管道
	Name string      //用户名
	Addr string
}

//保存在线用户cliAddr =====> Client
var onlineMap map[string]Client

var message = make(chan string)

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		//给当前客户端发送信息
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + " : " + msg
	return
}

//处理用户连接
func HandleConn(conn net.Conn) {
	defer conn.Close()
	//获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String()

	//创建一个结构体 默认　用户名和网络地址一样
	cli := Client{make(chan string), cliAddr, cliAddr}
	//把结构体添加到map
	onlineMap[cliAddr] = cli

	//新开一个　ｇｏｒｏｕｔｉｎｅ　专门给客户端发送信息
	go WriteMsgToClient(cli, conn)
	//广播某个人在线
	//message <- "[" + cli.Addr + "]" + cli.Name + " : login"
	message <- MakeMsg(cli, "login")
	//提示　我是谁
	cli.C <- MakeMsg(cli, " I am here ")

	//对方是否主动退出
	isQuit := make(chan bool)
	hasData := make(chan bool) //对方是否有数据发送

	//新开一个　ｇｏｒｏｕｔｉｎｅ　接收用户发过来的数据
	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				//对方断开　或者出问题
				fmt.Println("conn.Read err = ", err)
				return
			}
			msg := string(buf[:n-1]) //通过　nc测试　多一个换行

			if len(msg) == 3 && msg == "who" {
				//遍历map，给当前用户发送所有成员
				conn.Write([]byte("user list: \n"))
				for _, tmp := range onlineMap {
					msg = tmp.Addr + " : " + tmp.Name + " \n"
					conn.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				//rename |mike
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				onlineMap[cliAddr] = cli
				conn.Write([]byte("rename ok \n"))
			} else {
				//转发此内容
				message <- MakeMsg(cli, msg)
			}

			hasData <- true //代表有数据

		}

	}()
	for {
		//通过select 检测channel的流动
		select {
		case <-isQuit:
			//当前用户从map移除
			delete(onlineMap, cliAddr)
			message <- MakeMsg(cli, " logout ") //广播谁下线了
			return
		case <-hasData:

		case <-time.After(30 * time.Second): //60s后超时了
			delete(onlineMap, cliAddr)
			message <- MakeMsg(cli, " time out leave out ") //广播谁下线了
			return
		}
	}
}

func Manager() {
	//给map　分派空间
	onlineMap = make(map[string]Client)
	for {
		msg := <-message //没有消息前　这里会阻塞
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

func main() {
	//监听
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net.listen err = ", err)
		return
	}

	defer listener.Close()

	//新开一个　goroutine 转发消息　只要有消息来了　遍历map 给map每个成员都发送次消息
	go Manager()

	// 循环阻塞等待用户连接
	for {
		conn, err1 := listener.Accept()
		if err1 != nil {
			fmt.Println(" listener.accept err1 = ", err1)
			continue
		}
		go HandleConn(conn) //处理用户连接
	}

	fmt.Println("Hello World!")
}
