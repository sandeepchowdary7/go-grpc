package main

import (
	"context"
	pb "github.com/sandeepchowdary7/go-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListAllBlogs(in *emptypb.Empty, stream pb.BlogService_ListAllBlogsServer) error {
	cur, err := collection.Find(context.Background(), primitive.D{})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			"Internal Error")
	}

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				"Decode Err")
		}
		stream.Send(DocumentToBlog(data))
	}
	return nil
}
