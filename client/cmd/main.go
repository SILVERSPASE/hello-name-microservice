package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/micro/go-micro"
	"github.com/silverspase/hello-name-microservice/proto"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("proto.client"))
	service.Init()
	// Create new proto client
	greeter := proto.NewGreeterService("proto", service.Client())
	handler := myHandler{greeter}
	log.Print("starting server")
	log.Fatal(http.ListenAndServe(":8080", handler))
}



type Message struct {
	Name string `json:"name"`
}

type myHandler struct {
	greeter proto.GreeterService
}

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := Message{Name: "Wanderer"} // set default name

	err := json.NewDecoder(r.Body).Decode(&msg)
	switch {
	case err == io.EOF: // if body is empty, just go ahead
		break
	case err != nil:
		log.Print(err)
		_, err = fmt.Fprintf(w, fmt.Sprintf("%s\n", err))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// Call the proto
	rsp, err := h.greeter.Hello(context.TODO(), &proto.HelloRequest{Name: msg.Name})
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(w, fmt.Sprintf("%s\n", rsp.Greeting))
	if err != nil {
		log.Fatal(err)
	}
}