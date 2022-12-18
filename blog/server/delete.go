package main

import (
	"context"
	pb "github.com/sandeepchowdary7/go-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot Parse ID",
		)
	}
	res, err := collection.DeleteOne(ctx, bson.M{"_oid": oid})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot Delete Object in Mongo",
		)
	}
	if res.DeletedCount == 0 {
		return nil, status.Error(codes.NotFound, "Blog was not found")
	}
	return &emptypb.Empty{}, nil
}
