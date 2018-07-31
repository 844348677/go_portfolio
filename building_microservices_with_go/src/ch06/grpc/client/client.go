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
	// grpc.Dial signature
	// func Dial(target string, opts ...DialOption) (*ClientConn, error) {
	//		return DialContext(context.Background(), target, opts...)
	//	}
	conn , err := grpc.Dial("127.0.0.1:9000",grpc.WithInsecure())
	//  the target is a string which corresponds to the servver's network location and port
	// opts is a variadic list of DialOptions
	// WithInsecure , which disables transport security for the client ;  the default is that transport security is set so

	// the list of choices is very comprehensive , and you can specify configuration such as timeouts and using a load balancer
	// for the full list , documents https://godoc.org/google.golang.org/grpc#WithInsecure

	if err != nil{
		log.Fatalf("Unable to create connection to server : ",err)
	}
	// we are creating our client .
	// this is a type which is defined in our auto-generated code file
	// we pass it the connection we created ealier and then we can call the methods on the server
	client := proto.NewKittensClient(conn)
	response, err := client.Hello(context.Background(),&proto.Request{Name:"liuh"})

	if err != nil{
		log.Fatal("Error calling service : ",err)
	}

	fmt.Println(response.Msg)


}
