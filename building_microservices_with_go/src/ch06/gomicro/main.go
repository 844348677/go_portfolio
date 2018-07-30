package gomicro

import (

	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
	"log"
	"golang.org/x/net/context"
	kittens "./proto"
)
// creating a handler is as simple as defining a struct and adding methods to it
// Micro will automatically register these as routes on your service
type Kittens struct{}

func (s*Kittens) Hello(ctx context.Context, req *kittens.Request, resp *kittens.Response) error{
	resp.Msg = server.DefaultId + ": hello " + req.Name

	return nil
}


func main(){
	// 都是 go-micro 的包

	cmd.Init()

	// we are initializing the micro server , passing it some options
	// pass our basic HTTP server an address to configure what the IP and port the server would bind to
	server.Init(
		server.Name("bmigo.micro.Kittens"),
		server.Version("1.0.0"),
		server.Address(":8091"),
	)

	// Micro uses exactly the same signature which is present in the net/rpc package
	//Register Handlers
	server.Handle(
		server.NewHandler(
			// Kittens 结构 !!!!!
			new(Kittens),
		),
	)

	// Run server
	if err := server.Run();err != nil{
		log.Fatal(err)
	}

}


