package main

import (
	"log"
	"net/http"

	"github.com/micro/go-micro"
	"github.com/silverspase/go-micro/proto"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("proto.client"))
	service.Init()
	// Create new proto client
	greeter := proto.NewGreeterService("proto", service.Client())
	handler := handler{greeter}
	log.Print("starting server")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
