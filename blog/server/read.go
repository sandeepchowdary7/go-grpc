package main

import (
	"context"
	pb "github.com/sandeepchowdary7/go-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot Parse Mongo ID",
		)
	}

	data := &BlogItem{}
	res := collection.FindOne(ctx, bson.M{"_oid": oid})

	if err = res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.Internal, "Cannot Find Blog with this ID")
	}

	return DocumentToBlog(data), nil
}
