package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kevinwmiller/digidoc/grpc"
	"google.golang.org/grpc"
)

const (
	address = "localhost:3000"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDigiDocClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Upload(ctx, &pb.UploadRequest{
		Data: []byte("abcdefghijklmnopqrstuvwxyz"),
		MetaData: &pb.UploadMetaData{
			Name:       "NewDocument",
			ParentPath: "./",
			Tags:       []string{"first, second"},
		},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Uuid: %d", r.Uuid)
}
