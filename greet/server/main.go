package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/nithinkm8055/grpc-go-course/greet/proto"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lsnr, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v\n", addr, err)
	}

	log.Printf("listening on %s\n", addr)

	s := grpc.NewServer()

	// register the service implementation on the grpc server
	pb.RegisterGreetServiceServer(s, &Server{})

	err = s.Serve(lsnr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v\n", addr, err)
	}
}
