package main

import (
	"context"
	pb "github.com/sandeepchowdary7/go-grpc/greet/proto"
	"log"
)

func doGreet(client pb.GreetServiceClient) {

	log.Printf("Client was Invoked")

	res, err := client.Name(context.Background(), &pb.GreetRequest{
		FirstName: "Sandeep",
		LastName:  "Chowdary",
	})

	if err != nil {
		log.Fatalf("something wrong with client %s/n", err)
	}

	log.Printf("Greet : %s", res)
}
