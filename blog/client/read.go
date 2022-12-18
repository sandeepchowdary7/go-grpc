package main

import (
	"context"
	pb "github.com/sandeepchowdary7/go-grpc/blog/proto"
	"log"
)

func readBlog(c pb.BlogServiceClient, id string) {
	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Error Happening while Reading: %v/n", err)
	}

	log.Printf("Blog was Read %v/n", res)
}
