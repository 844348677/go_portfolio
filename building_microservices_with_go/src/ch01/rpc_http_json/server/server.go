package server

import (
	"../contract"
	"net/rpc"
	"net"
	"fmt"
	"log"
	"net/http"
	"net/rpc/jsonrpc"
	"io"
)



const(
	DefaultRPCPath = "/_goRPC_"
	DefaultDebugPath = "/debug/rpc"
)


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

	// main different is here
	// instead of starting the RPC server we are starting an http server and passing the listener to it along with a handler
	http.Serve(listen,http.HandlerFunc(httpHandler))
}
// the handler we are passing to the server is where the magic happens
func httpHandler(w http.ResponseWriter,r *http.Request){
	// calling this function
	// passing to it a type that implements io.ReadWriteCloser
	// return a type that implements rpc.ClientCodec
	serverCodec := jsonrpc.NewServerCodec(&HttpConn{in:r.Body,out:w})
	err := rpc.ServeRequest(serverCodec)
	if err != nil{
		log.Printf("Error while serving JSON request : %v", err)
		http.Error(w,"Error while serving Josn request , details have ben logge.",500)
		return
	}

}

// 又没讲这个　struct ! 好吧　讲了　在最后面
type HttpConn struct {
	in  io.Reader
	out io.Writer
}

func (c *HttpConn) Read(p []byte) (n int, err error)  { return c.in.Read(p) }
func (c *HttpConn) Write(d []byte) (n int, err error) { return c.out.Write(d) }
func (c *HttpConn) Close() error                      { return nil }


type HelloWorldHandler struct{}

// import contract 费了半天劲
func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply * contract.HelloWorldResponse) error{
	reply.Message = "Hello " + args.Name
	return nil
}