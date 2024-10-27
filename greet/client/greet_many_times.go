package main

import (
	"context"
	"io"
	"log"

	pb "github.com/nithinkm8055/grpc-go-course/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) error {
	stream, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{
		FirstName: "Hulkbuster",
	})
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("cannot receive stream: %v", err)
		}

		log.Printf("Received Greeting: %s", resp.Result)
	}

	return nil
}
