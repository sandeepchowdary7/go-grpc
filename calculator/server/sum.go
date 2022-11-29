package main

import (
	"context"

	pb "github.com/sandeepchowdary7/go-grpc/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{
		Z: in.X + in.Y,
	}, nil
}
