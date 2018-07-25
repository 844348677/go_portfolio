package main

import (
	"net/http"
	"encoding/json"
)

type helloWorldResponse struct{
	Message string `json:"message"`
}

// send our data to the output stream without marshalling to a temporary object before we return it

func helloWorldHandler(w http.ResponseWriter, r *http.Request){
	response := helloWorldResponse{Message:"hello - world"}
	// func NewEncoder(w io.Writer) *Encoder
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
	// so instead of storing the output of marshal into a byte array, we can write it straight to the http
}

