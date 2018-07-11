// 04_httpserver
package main

import (
	"fmt"
	"net/http"
)

// w 给客户端回复数据
//　req　读取客户端发送的数据
func HandleConn(w http.ResponseWriter, req *http.Request) {
	// 请求行中的　请求方法　get post
	fmt.Println("req.Method = ", req.Method)
	fmt.Println("req.URL = ", req.URL)
	fmt.Println("req.header = ", req.Header)
	fmt.Println("req.body = ", req.Body)

	w.Write([]byte("hello go")) //给客户端回复数据

}

func main() {
	//注册　处理函数　用户连接　自动调用指定的处理函数
	http.HandleFunc("/go", HandleConn)

	//监听绑定
	http.ListenAndServe(":8000", nil)
	fmt.Println("Hello World!")
}
