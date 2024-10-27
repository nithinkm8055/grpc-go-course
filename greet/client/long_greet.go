package main

import (
	"context"
	"log"

	pb "github.com/nithinkm8055/grpc-go-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) error {
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		return nil
	}

	reqs := []*pb.GreetRequest{
		{
			FirstName: "Thor",
		},
		{
			FirstName: "Loki",
		},
		{
			FirstName: "Odin",
		},
	}

	for _, req := range reqs {
		if err := stream.Send(req); err != nil {
			return err
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Printf("Received Greeting: %s\n", resp.Result)
	return nil
}
