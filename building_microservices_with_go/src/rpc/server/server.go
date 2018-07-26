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
	// first creata a new instance of our handler
	// then we register this with the default RPC server
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)

	port := "8080"
	// 这里书写错了　搞了那么多引号　""
	// create a socket using the given protocol and binding it to the IP address and port
	// tcp tcp4 tcp6 unix unixpacket
	// Listen() function returns an instance that implements the Listener interface
	// type Listener interface{ ... } 里面有三个方法
	listen,err := net.Listen("tcp",fmt.Sprintf(":%v",port))
	if err != nil{
		log.Fatal(fmt.Sprintf("unable to listen on given prot : %s",err))
	}

	// endless for loop
	// handle each connection individually
	for {
		conn,_ := listen.Accept()
		// func ServeConn(conn io.ReadWriteCloser)
		// gob
		// the source and destination values and types do not need to correspond exactly
		// if a field is in the source but not in the receiving struct
		// then the decoder will ignore this field and the processing will continue without errors
		// JMI : the exact same interface must be present on both the client and server
		go rpc.ServeConn(conn)
	}
	// to make a request
	// no longer simply use curl
	// no longer using the HTTP protocol and the message format is no longer JSON


}





