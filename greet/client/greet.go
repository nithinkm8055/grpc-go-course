package main

import (
	"context"
	"log"

	pb "github.com/nithinkm8055/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) error {
	resp, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Hulkbuster",
	})
	if err != nil {
		return err
	}

	log.Printf("Greeting %s\n", resp.Result)
	return nil
}
