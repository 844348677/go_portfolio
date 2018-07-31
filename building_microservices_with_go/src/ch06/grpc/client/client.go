package main

import (
	"google.golang.org/grpc"
	"log"
	"../proto"
	"context"
	"fmt"
)

// creating a connection to our server and then initiating the request

func main(){
	conn , err := grpc.Dial("127.0.0.1:9000",grpc.WithInsecure())

	if err != nil{
		log.Fatalf("Unable to create connection to server : ",err)
	}
	client := proto.NewKittensClient(conn)
	response, err := client.Hello(context.Background(),&proto.Request{Name:"liuh"})

	if err != nil{
		log.Fatal("Error calling service : ",err)
	}

	fmt.Println(response.Msg)


}
