package main

import (
	"github.com/koding/kite"
	"github.com/koding/kite/config"
	"net/url"
	"fmt"
)

// Setup
// go get github.com/koding/kite
// there is a service which runs along with your application Kites called kontrol
// this handles service discovery, and all of your application services register with this service so that the clients can query the service catalog to obtain the service endpoit
// etcd

func main() {
	// we are creating our new Kite
	// two arguments : the name of our Kite and  the service version
	// obtain a reference to the configuration and set this to our Kite
	k := kite.New("math","1.0.0")

	c := config.MustGet()
	k.Config = c
	// to be able to register our kite with the server discovery , we have to set the KontrolURL with the correct URI fro our kontrol server
	k.Config.KontrolURL = "http://kontrol:6000/kite"
	// we are using the name that is supplied by Docker when we link together some containers

	// we are registering our container with the kontrol server.
	// we need to pass the URL scheme we are usings
	// in this instance , HTTP is the hostname; this needs to the accessible name for the application
	k.RegisterForever(&url.URL{Scheme:"http",Host:"127.0.0.1:8091",Path:"/kite"})



	//add our handler method with the name "square"
	k.HandleFunc("Hello",func(r *kite.Request)(interface{},error){
		// the signature for HandleFunc is very similar to that of the standard HTTP library;
		// we set up a route and pass a fucntion which would be responsible fro executing that request
		name,_ := r.Args.One().String()

		// dnode

		return fmt.Sprintf("Hello %v",name),nil
	}).DisableAuthentication()

	//Attach to a server with port8091 and run it
	k.Config.Port = 8091
	k.Run()

	// One of the nice features of Kites is that authentication is built in and ,
	// it is quite common to restrict the actions of a particular service call based upon the permissions of the caller
	//User the hood , Kite is using JWT to break down these permissions into a set of claims
	// the principle is that a key is signed and therefore a receiving service only has to validate the signature of the key to trust its payload rather than having to call a downstream service

}

// Code generation
// there is no code generation or templates

// Tooling
// etcd , which is used for your service discovery
// etcd and kite are easily packaged into a Docker container

// Maintainable

// Format
// Kite uses dnode as its messaging protocol

// Patterns
// Service discovery is built into Kite with the application kontrol
// The backend store for kontrol is not propriety, but it uses a plugin architecture and supports  etcd,consul, and son