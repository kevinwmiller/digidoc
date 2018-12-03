package client

import (
	"context"

	pb "github.com/kevinwmiller/digidoc/grpc"
	"google.golang.org/grpc"
)

type UploadConfig struct {
	ServerAddr string,
	Filename string,
}

func Upload(cfg *UploadConfig) error {
	c := &DigidocGrpcUploadClient {
		ServerAddr: cfg.ServerAddr,
		Filename: cfg.Filename,
	}
	c.Upload(context.Background(), cfg)
}
