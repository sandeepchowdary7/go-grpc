package main

import (
	"context"
	pb "github.com/sandeepchowdary7/go-grpc/blog/proto"
	"log"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	newBlog := pb.Blog{
		Id:       id,
		AuthorId: "Sandeep77",
		Title:    "Test77",
		Content:  "Test Content77",
	}
	_, err := c.UpdateBlog(context.Background(), &newBlog)

	if err != nil {
		log.Fatalf("Blog Was Not Updated: %v/n", err)
	}
	log.Printf("Blog Was Updated Succesfully")
}
