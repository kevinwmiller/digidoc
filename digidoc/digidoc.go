package digidoc

import (
	"context"
	"io"
	"fmt"

	"github.com/pkg/errors"
	pb "github.com/kevinwmiller/digidoc/grpc"
)

type Digidoc struct {
	// Database connection
	// config
}

func (d *Digidoc) Upload(stream pb.DigiDoc_UploadServer) (error) {
	fmt.Println("Uploading data")
	for {
		data, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				goto END
			}

			err = errors.Wrapf(err,
				"failed unexpectadely while reading chunks from stream")
			return err
		}
		fmt.Println("Received data ", data)
	}

END:
	err := stream.SendAndClose(&pb.UploadStatus{
		Message: "Upload received with success",
		Code:    pb.UploadStatusCode_Ok,
	})
	return err
}

func (d *Digidoc) Browse(ctx context.Context, in *pb.BrowseRequest) (*pb.BrowseResponse, error) {
	return &pb.BrowseResponse{}, nil
}
