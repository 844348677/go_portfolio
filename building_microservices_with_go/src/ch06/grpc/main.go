package main

import "context"
import (
	"./proto"
	"fmt"
	"log"
	"google.golang.org/grpc"
	"net"
)

type kittenServer struct {}

func (k *kittenServer) Hello(ctx context.Context, request *proto.Request) (*proto.Response,error){
	response := &proto.Response{}
	response.Msg = fmt.Sprintf("Hello %v",request.Name)
	return response,nil
}

// KittensServer is the server API for Kittens service.
/*
type KittensServer interface {
	Hello(context.Context, *Request) (*Response, error)
}
*/
// we have our context and an object which corresponds to the request message we defined in our protos file and return tuple of response and error

func main(){
	lis,err := net.Listen("tcp",fmt.Sprintf(":%d",9000))
	if err != nil{
		log.Fatalf("failed to listen : %v",err)
	}
	grpcServer := grpc.NewServer()

	proto.RegisterKittensServer(grpcServer,&kittenServer{})

	grpcServer.Serve(lis)
}

