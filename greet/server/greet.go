package main

import (
	"context"
	pb "github.com/sandeepchowdary7/go-grpc/greet/proto"
)

func (s *Server) Name(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		FullName: in.FirstName + in.LastName,
	}, nil
}
