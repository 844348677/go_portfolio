package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"fmt"
)

// func Unmarshal(data []byte,v interface{}) error

// map[string]interface{}

type helloWorldRequest struct{
	Name string `json:"name"`
}
// you do not write code for the compiler , you write code for humans to understand
// you will spend more time reading code than you do writing it
// {"name":"World"}
// {"Name":"World"}

// Method string
// Header Header
// Body io.ReadCloser

// the JSON that has been sent with the request is accessible in the Body fild
// io.ReadCloser


func helloWorldHandler(w http.ResponseWriter, r *http.Request){
	body,err := ioutil.ReadAll(r.Body)
	if err != nil{
		http.Error(w,"Bad request",http.StatusBadRequest)
		return
	}

	var request helloWorldRequest
	err = json.Unmarshal(body, & request)
	if err != nil{
		http.Error(w,"bad  request",http.StatusBadRequest)
		return
	}
	response := helloWorldResponse{Message:"Hello " + request.Name}
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
type helloWorldResponse struct{
	Message string `json:"message"`
}

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

// liuh@liuh-pc:~/go/go_portfolio/building_microservices_with_go/src/r_w_json_4$ go run ./json_4.go

// curl localhost:8080/helloworld -d '{"name":"LiuH"}'
// {"message":"Hello LiuH"}

// curl localhost:8080/helloworld
// func Error(w ResponseWriter,error string,code int)
// bad  request


