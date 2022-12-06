package main

import (
	"context"
	"fmt"
	pb "github.com/sandeepchowdary7/go-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	data := BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}
	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("internal error"),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("cannot convert OID"),
		)
	}
	return &pb.BlogId{Id: oid.Hex()}, nil
}
