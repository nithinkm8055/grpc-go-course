package main

import (
	"io"
	"log"

	pb "github.com/nithinkm8055/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) LongGreet(stream grpc.ClientStreamingServer[pb.GreetRequest, pb.GreetResponse]) error {
	log.Printf("LongGreet func was invoked")

	var res string = "Greeting Long Greet: "

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			return err
		}

		res += "Hello " + req.FirstName + "!\n"
	}
}
