package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/silverspase/go-micro/proto"
)

type Message struct {
	Name string `json:"name"`
}

type handler struct {
	greeter proto.GreeterService
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
