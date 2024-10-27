package main

import (
	"fmt"
	"log"

	pb "github.com/nithinkm8055/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) GreetManyTimes(req *pb.GreetRequest, stream grpc.ServerStreamingServer[pb.GreetResponse]) error {
	log.Printf("GreetManyTimes func invoked for %s", req.FirstName)
	firstName := req.FirstName

	for i := range 10 {
		if err := stream.Send(&pb.GreetResponse{
			Result: fmt.Sprintf("Greeting Many Times: Hi %s, number %d", firstName, i),
		}); err != nil {
			return err
		}
	}

	return nil
}
