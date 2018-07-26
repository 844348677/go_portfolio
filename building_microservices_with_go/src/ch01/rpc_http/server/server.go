package server

import (
	"../contract"
	"net/rpc"
	"net"
	"fmt"
	"log"
	"net/http"
)


// use HTTP as transport protocol then the rpc package can facility this by calling the HandleHTTP method

// HandleHTTP method set up two endpoint in

const(
	DefaultRPCPath = "/_goRPC_"
	DefaultDebugPath = "/debug/rpc"
)

// DefaultDebugPath : detail fro the registered endpoint
// two thing
// this does not mean you can communicate easily with API from a web browser
// the message are still gob encoded so you would need to write a gob encoder and decoder in JavaScript
//


func StartServer(){
	helloWorld:= &HelloWorldHandler{}
	rpc.Register(helloWorld)
	// this is a requiremeng using HTTP with RPC as it will register the HTTP handlers
	rpc.HandleHTTP()

	port := "8080"
	listen,err := net.Listen("tcp",fmt.Sprintf(":%v",port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port : ",err))
	}
	log.Printf("server starting on port %v\n",port)
	// nil DefaultServer
	http.Serve(listen,nil)
}

type HelloWorldHandler struct{}

// import contract 费了半天劲
func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply * contract.HelloWorldResponse) error{
	reply.Message = "Hello " + args.Name
	return nil
}