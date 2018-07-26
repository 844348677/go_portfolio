package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"log"
)

type helloWorldResponse struct{
	Message string `json:"message"`
}

func helloWorldHandler(w http.ResponseWriter,r *http.Request){
	response := helloWorldResponse{Message:"Hello wold"}
	data,err := json.Marshal(response)
	if err != nil{
		panic("Ooops")
	}
	// we are checking to see if there is a callback parameter in the query sting
	// this would be provided by the client and indicates the function they expect to be called when the response is returned

	callback := r.URL.Query().Get("callback")
	if callback != ""{
		// 书上代码有问题 github　上的代码也有问题　改请求头有啥用？　应该是响应头
		//r.Header.Add("Content-Type","application/javascript")
		w.Header().Add("Content-Type","application/javascript")
		fmt.Fprintf(w,"%s (%s)",callback,string(data))
	}else{
		fmt.Fprint(w,string(data))
	}

}

// Request
// http://localhost:8080/helloworld?callback=hello
// http://localhost:8080/helloworld

func main(){
	port := 8080
	http.HandleFunc("/helloworld",helloWorldHandler)
	log.Printf("server staring on port %v\n",8080)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}
