package main

import (
	"net/http"
	"encoding/json"
	"log"
	"fmt"
)

// 性能问题　
// Encoder Decoder 减少中间变量，减少代码　直接使用流处理

func helloWorldHandler(w http.ResponseWriter,r *http.Request){
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil{
		http.Error(w,"Bad request " ,http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message:"Hello " + request.Name}
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

type helloWorldRequest struct{
	Name string `json:"name"`
}
type helloWorldResponse struct{
	Message string `json:"message"`
}

func main(){
	port := 8080
	http.HandleFunc("/helloworld",helloWorldHandler)
	log.Printf("server staring on port %v\n",8080)

	daghandler := http.FileServer(http.Dir("./images"))
	http.Handle("/dog/",http.StripPrefix("/dog/",daghandler))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}

// http://localhost:8080/dog/1.jpeg

// if we did not use StripPrefix , then the FileServer handler would be looking for our images in the immages/dog directory


