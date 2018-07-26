package server

import (
	"../contract"
	"net/rpc"
	"net"
	"fmt"
	"log"
)

// the Go standard library has fantastic support for RPC right out-of-the-box

// standard RPC
// client an server
// standard interface to communicate over RPC
// https://en.wikipedia.org/wiki/Remote_procedure_call

// contract是毛线　非要去github上找源码！　骗星行为！

type HelloWorldHandler struct{}

// import contract 费了半天劲
func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply * contract.HelloWorldResponse) error{
	reply.Message = "Hello " + args.Name
	return nil
}

// we alse define a handler
// difference  handler and http.handler
// it does not need to conform to an interface
// we have a struct field with method on it
// we can register this with the RPC server
// func Register(rcvr interface{}) error

// if my client wanted to cal the HelloWorld method
// we would access it using HelloWorldHandler.HelloWorld

// func ResiterName(name string,rcvr interface{}) error

func StartServer(){
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)

	port := "8080"
	// 这里书写错了　搞了那么多引号　""
	listen,err := net.Listen("tcp",fmt.Sprintf(":%v",port))
	if err != nil{
		log.Fatal(fmt.Sprintf("unable to listen on given prot : %s",err))
	}

	for {
		conn,_ := listen.Accept()
		go rpc.ServeConn(conn)
	}
	// first creata a new instance of our handler
	// then we register this with the default RPC server
}





