package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/nithinkm8055/grpc-go-course/greet/proto"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to %s: %v\n", addr, err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	if err := doGreet(c); err != nil {
		log.Printf("cannot receive greet: %v\n", err)
		return
	}

	if err := doGreetManyTimes((c)); err != nil {
		log.Printf("cannot receive greet many times: %v\n", err)
		return
	}

	if err := doLongGreet((c)); err != nil {
		log.Printf("cannot receive long greet: %v\n", err)
		return
	}

	if err := doGreetEveryone((c)); err != nil {
		log.Printf("cannot receive greet everyone: %v\n", err)
		return
	}
}
