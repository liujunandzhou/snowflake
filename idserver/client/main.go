package main

import (
	"log"

	pb "github.com/liujunandzhou/snowflake/idserver/idserver"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewIdServerClient(conn)

	// Contact the server and print out its response.

	for i := 0; i < 10000; i++ {

		r, err := c.GetId(context.Background(), &pb.Request{})

		if err != nil {

			log.Fatalf("could not greet: %v", err)

		}

		log.Printf("UniqId: %s", *r.Uniqid)

	}
}
