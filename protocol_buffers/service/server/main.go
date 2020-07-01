package main

import (
	"github.com/ByronLeeLee/go/study/protocol_buffers/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("fail to listen")
	}

	s := grpc.NewServer()

	service.RegisterGreeterServer(s, &service.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve")
	}
}
