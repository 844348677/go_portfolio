// 16_file
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteFile(path string) {
	//打开文件
	//使用完毕，需要关闭文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//使用完毕　需要关闭文件
	defer f.Close()

	var buf string

	for i := 0; i < 10; i++ {
		// “i=1\n", 字符串存储在buf中
		buf = fmt.Sprintf("i = %d\n", i)
		n, err := f.WriteString(buf)
		if err != nil {
			fmt.Println("error = ", err)
		}
		fmt.Println(" n = ", n)
	}
}
func ReadFile(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//关闭文件
	defer f.Close()

	buf := make([]byte, 1024*2) //2k大小
	// n 代表文件读取内容的长度
	n, err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF { //文件出错，同时没到结尾　io.EOF表示文件结尾
		fmt.Println("err1 = ", err1)
		return
	}
	fmt.Println("buf = ", string(buf[:n]))

}

//每次读取一行
func ReadFileLine(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//关闭文件
	defer f.Close()

	//新建一个缓冲区，把内容先放在缓冲区
	r := bufio.NewReader(f)
	//遇到　'\n'　结束读取 但是　'\n'也读取进去了
	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			//文件已经结束
			if err == io.EOF {
				break
			}
			fmt.Println("err = ", err)
		}

		fmt.Printf(" buf = #%s# \n", string(buf))
	}

}
func main() {
	path := "./demo.txt"
	//WriteFile(path)

	//ReadFile(path)
	ReadFileLine(path)
	fmt.Println("Hello World!")
}
