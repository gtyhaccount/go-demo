package main

import (
	"context"
	"github.com/ByronLeeLee/go/study/protocol_buffers/service"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const address = "localhost:50051"
const defaultName = "world"

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("get conn fail")
	}

	defer conn.Close()
	client := service.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.SayHello(ctx, &service.Request{Name: name})
	if err != nil {
		log.Fatalf("could not get:%v", err)
	}

	log.Printf("Getting: %s", response)
}
