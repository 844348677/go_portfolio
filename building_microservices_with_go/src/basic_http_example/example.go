package main

import (
	"net/http"
	"log"
	"fmt"
)
// introduction to microservices
// first ,create a simple web server using net/http

// building a simple web server with net/http
// single endpoint
// Transport Layer Security TLS

func main(){
	port := 8080
	http.HandleFunc("/helloworld",helloWorldHandler)
	log.Printf("server staring on port %v\n",8080)

	// start the http server
	// first paramter port
	// second parameter is nil , using DdefaultServeMux
	// later
	// bind to a port that is already in use on the server
	// log.Fatal(err)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}

// the first thing we are doing is calling the HandlerFunc
// creates a Handler
// mapping the path   (first parameter)

func helloWorldHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"hello world\n")
}

// ps -aux | grep 'go run'
// http://127.0.0.1:8080/helloworld
