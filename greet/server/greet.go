package main

import (
	"context"
	"log"

	pb "github.com/nithinkm8055/grpc-go-course/greet/proto"
)

func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet func invoked for %s", req.FirstName)

	return &pb.GreetResponse{
		Result: "Hello " + req.FirstName + "👋",
	}, nil
}
