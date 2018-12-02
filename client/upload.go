package client

import (
	"context"

	pb "github.com/kevinwmiller/digidoc/grpc"
	"google.golang.org/grpc"
)

func (c *Client) Upload(ctx context.Context, opts ...grpc.CallOption) (pb.DigiDoc_UploadClient, error) {
	
}
