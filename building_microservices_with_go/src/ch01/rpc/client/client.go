package client

import (
	"net/rpc"
	"fmt"
	"log"
	"../contract"
)

// show how we need to setup rpc.client
func CreateClient() *rpc.Client{
	port := "8080"
	// func Dial(network,address string) (*Client,error)
	client,err := rpc.Dial("tcp",fmt.Sprintf("localhost:%v",port))
	if err != nil {
		log.Fatal("dialing: ",err)
	}
	return client
}

// we then use this returned connection to make a request to the server

func PerformRequest(client *rpc.Client) contract.HelloWorldResponse{
	args := &contract.HelloWorldRequest{Name:" World!"}
	var reply contract.HelloWorldResponse

	err := client.Call("HelloWorldHandler.HelloWorld",args,&reply)
	if err != nil{
		log.Fatal("error: ",err)
	}
	return reply
}
// Call : waits util the server sends a reply writing the response assuming there is no error to the reference of
// our HelloWorldResponse passed to the method and