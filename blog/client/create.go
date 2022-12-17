package main

import (
	"context"
	pb "github.com/sandeepchowdary7/go-grpc/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	blog := &pb.Blog{
		AuthorId: "Sandeep",
		Title:    "Test",
		Content:  "Test Content",
	}
	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("unexpected err %s/n", err)
	}
	log.Printf("Blog has been Created: %s/n", res)

	return res.Id
}
