package ch04

import (
	"net/http"
	"log"
)

func main(){
	err:= http.ListenAndServe(";2323",&handler.SearchHander{})
	if err != nil{
		log.Fatal(err)
	}
}

// 这章没有必要打代码
