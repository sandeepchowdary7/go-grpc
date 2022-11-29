package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/sandeepchowdary7/go-grpc/greet/proto"
)

var addr = "0.0.0.0.50551"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Connection Failed %s/n", err)
	}

	log.Printf("Listening on %v/n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve %s/n", err)
	}
}
