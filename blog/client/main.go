package main

import (
	"google.golang.org/grpc/credentials/insecure"
	"log"

	pb "github.com/sandeepchowdary7/go-grpc/blog/proto"
	"google.golang.org/grpc"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to Connect %v/n", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Defer Function Error %v/n", err)
		}
	}(conn)

	c := pb.NewBlogServiceClient(conn)
	id := createBlog(c)
	deleteBlog(c, id)
	updateBlog(c, id)
	readBlog(c, id)
	listAllBlogs(c)
}
