package main

import (
	"context"
	"io"
	"log"

	pb "github.com/nithinkm8055/grpc-go-course/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) error {
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		return err
	}

	reqs := []*pb.GreetRequest{
		{
			FirstName: "Stark",
		},
		{
			FirstName: "Hogan",
		},
		{
			FirstName: "Potts",
		},
	}

	go func() error {
		for _, req := range reqs {
			err := stream.Send(req)
			if err != nil {
				return err
			}
		}
		stream.CloseSend()
		return nil
	}()

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("Received Greeting: %s\n", resp.Result)
	}

	return nil
}
