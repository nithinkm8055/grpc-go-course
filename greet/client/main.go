package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/nithinkm8055/grpc-go-course/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	caCertfile := "ssl/ca.crt"
	creds, err := credentials.NewClientTLSFromFile(caCertfile, "")
	if err != nil {
		log.Fatalf("error loading ca trust certificate %v\n", err)
	}
	opts := []grpc.DialOption{}

	opts = append(opts, grpc.WithTransportCredentials(creds))

	conn, err := grpc.NewClient(addr, opts...)
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
