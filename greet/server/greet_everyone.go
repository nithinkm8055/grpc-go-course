package main

import (
	"io"
	"log"

	pb "github.com/nithinkm8055/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) GreetEveryone(stream grpc.BidiStreamingServer[pb.GreetRequest, pb.GreetResponse]) error {
	log.Printf("GreetEveryone func was invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if err := stream.Send(&pb.GreetResponse{
			Result: "Greeting GreetEveryone: Hi " + req.FirstName + "!\n",
		}); err != nil {
			return err
		}
	}

	return nil
}
