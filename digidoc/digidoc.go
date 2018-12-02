package digidoc

import (
	"context"
	"fmt"

	pb "github.com/kevinwmiller/digidoc/grpc"
)

type Digidoc struct {
	// Database connection
	// config
}

func (d *Digidoc) Upload(ctx context.Context, in *pb.UploadRequest) (*pb.UploadResponse, error) {
	fmt.Println(in.MetaData.Name)
	return &pb.UploadResponse{Uuid: 52}, nil
}

func (d *Digidoc) Browse(ctx context.Context, in *pb.BrowseRequest) (*pb.BrowseResponse, error) {
	return &pb.BrowseResponse{}, nil
}
