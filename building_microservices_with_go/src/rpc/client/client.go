package client

import (
	"net/rpc"
	"fmt"
	"log"
)

func CreateClient() *rpc.Client{
	port := "8080"
	client,err := rpc.Dial("tcp",fmt.Sprintf("localhost:%v",port))
	if err != nil {
		log.Fatal("dialing: ",err)
	}
	return client
}