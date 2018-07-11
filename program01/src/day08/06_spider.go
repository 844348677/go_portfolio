// 06_spider
package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

//爬取网页内容
func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	//读取网页内容
	buf := make([]byte, 1024*4)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			//读取结束　ＥＯＦ　或出错
			fmt.Println("resp.body.read err = ", err)
			break
		}
		result += string(buf[:n])
	}
	return
}

func DoWork(start, end int) {
	fmt.Printf("正在爬去　%d 到　%d　的页面\n", start, end)

	//明确目标
	//https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=0
	//https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=50
	//https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=100

	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		fmt.Println(" url = ", url)

		//2 爬　将所有的网站的内容　全部爬出来
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println("httpGet err = ", err)
			continue
		}
		//把内容写入到文件
		fileName := strconv.Itoa(i) + ".html"
		f, err1 := os.Create(fileName)
		if err1 != nil {
			fmt.Println("os.create err1 = ", err1)
			continue
		}
		//写内容
		f.WriteString(result)
		//关闭文件
		f.Close()
	}
}

func main() {
	var start, end int
	fmt.Printf("请输入起始页　( >=1): ")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页　( >=起始页): ")
	fmt.Scan(&end)

	DoWork(start, end)

	fmt.Println("Hello World!")
}
