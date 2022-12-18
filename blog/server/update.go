package main

import (
	"context"
	"fmt"
	pb "github.com/sandeepchowdary7/go-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Internal Error",
		)
	}

	data := &BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}
	res, err := collection.UpdateOne(ctx, bson.M{"-id": oid}, bson.M{"$set": data})

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Blog Not Updated",
		)
	}
	if res.MatchedCount == 0 {
		fmt.Println("Cannot Find Blog ID")
	}
	return &emptypb.Empty{}, nil
}
