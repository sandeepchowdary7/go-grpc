package main

import (
	"context"
	pb "github.com/sandeepchowdary7/go-grpc/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
)

func listAllBlogs(c pb.BlogServiceClient) {
	stream, err := c.ListAllBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling Blog %v/n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		log.Println(res)
	}

}
