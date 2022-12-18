package main

import (
	"context"
	pb "github.com/sandeepchowdary7/go-grpc/blog/proto"
	"log"
)

func deleteBlog(c pb.BlogServiceClient, id string) {

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("Error While Deleting a Blog: %v/n", err)
	}
	log.Println("Blog Deleted Successfully")
}
