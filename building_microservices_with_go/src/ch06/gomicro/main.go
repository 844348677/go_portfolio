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
// the form of the handler looks very similar to the one from the net/http package
// we can see the same context object  we looked at in the first chapter
// the Context is a safe method fro accessing request scoped data which can be accessed from multiple Goroutines
// the request and the response objects are those which we defined in our proto file
// Instead of writing our response to a ResponseWriter in this handler,
// we set the values we wish to return to the reference of the response which is passed to the function



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

// Tooling (CI/CD, cross platform)
// because Micro is written in pure Go, with the only external dependency being protoc ,
// it creates a very lightweight framework which would be possible to use on Linux , Mac , and Windows with ease.
// It would alse be easy to set up onto a CI server;
// the main complexity is the installation of protoc , but this application is incredibly well supported by Google and is available for all the main operating systems and architectures

// Maintainable
// The way that Micro has been built is incredibly sympathetic towards modern enterprise problems of updating and maintaining microservices.
// Vesioning is incorporated into the framework,
// and in our example , we are setting the version in the server.Init method
// it is possible for multiple services to co-exit differentiated by their version number.
// When requesting the service , it is possible to filter by a version which allows new version of a servcie to be deployed without causing disruption to the rest of the estate

// Format (REST/RPC)
// Micro uses Google's Protocol Buffers as it core message protocol.
// this is not the only method by which you can communicate with the services
// micro also implements the sidecar pattern which is an RPC proxy
// This gives a really simple way of intergrating any application into the Micro ecosystem
// the sidecar can be used as an API gateway , which is a single point of entry fro multiple downstream services
// In Micro's case, the gateway handles HTTP request and converts them to RPC
// it is also capable of providing reverse-proxy functionality
// https://micro.mu/

// Patterns

// Language independence
// use of Protocol Buffers as a messaging format and the ability of the siecar
// it is possible to interface with microservices from just about any language that can support an HTTP transport
// send and receive message using Ruby client/client.rb
//

// Ability to interface with other frameworks
// it si possible to intergrete Micro with many different frameworks

// Efficiency

// Quality

// Open source

// Security

// Support

// Extensibility
