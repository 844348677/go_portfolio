package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type helloWorldResponse struct{
	Message string `json:"message"`
}

// using th struct field's tages , we can have greater control of how the output will look


func helloWorldHandler(w http.ResponseWriter, r *http.Request){
	response := helloWorldResponse{Message:"hello - world"}
	data,err := json.Marshal(response)
	if err!=nil{
		panic("OOOOOOps")
	}
	fmt.Println(w,string(data))
}

// we can use field tags to control the output even further
type helloWorldResponse2 struct{
	// change the output field to be "message"
	Message string `json:message`

	// do not output this field
	Author string `json:"-"`

	// do not output the field if the value is empty
	Date string `json:",omitempty"`

	// convert output to a string and rename "id"
	Id int `json:"id,string"`
}
// channel , complex types ,and fucntions cannot be encoded in JSON
// UnsupportedTypeError
// it also cannot represent cyclic data structures