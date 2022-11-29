package main

import (
	pb "github.com/sandeepchowdary7/go-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr = "localhost:50550"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Connection Failed %s/n", err)
	}

	log.Printf("Listening on %v/n", addr)

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	g := pb.NewGreetServiceClient(conn)

	doGreet(g)
}
