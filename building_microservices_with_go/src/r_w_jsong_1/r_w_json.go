package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

// to encode json data
// enconding/json -> Marshal
// func Marshal(v interface{}) ([]byte,error)
// unhandled errors are a bad thing
// in the instance that Marshal cannot create a JSON encoded byte array from the given object
//	runtime panic, this is captured , is returned to the caller

type helloWorldResponse struct{
	Message string
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request){
	response := helloWorldResponse{Message:"hello - world"}
	data,err := json.Marshal(response)
	if err!=nil{
		panic("OOOOOOps")
	}
	fmt.Println(w,string(data))
}
