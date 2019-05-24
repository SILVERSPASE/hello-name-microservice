package main

import (
	"context"
	"log"

	"github.com/micro/go-micro"
	"github.com/silverspase/go-micro/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("proto"))
	// Init will parse the command line flags.
	service.Init()
	// Register handler
	err := proto.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err != nil {
		log.Fatal(err)
	}

	// Run the greeter-service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
