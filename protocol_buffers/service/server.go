package service

import (
	"context"
	"log"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	UnimplementedGreeterServer
}

func (s Server) SayHello(ctx context.Context, in *Request) (*Response, error) {
	log.Printf("Received: %v", in.GetName())
	return &Response{Message: "Hello " + in.Name}, nil
}
