package grpc

import (
	"errors"

	"google.golang.org/grpc"
)

type DigidocGrpcUploadClient struct {
	*DigidocGrpcClient,
	Filename string,
	client DigiDocClient,
	chunkSize uint64,
}

type DigidocGrpcUploadClientConfig struct {
	Address string,
	Filename string,
	ChunkSize uint64,
}

func NewDigidocGrpcUploadClient(cfg DigidocGrpcUploadClientConfig) (c DigidocGrpcUploadClient, err error) {
	grpcOpts := []grpc.DialOption{}
	if cfg.Address == "" {
		// TODO: This should look at DIGIDOC_SERVER_ADDR if not set. Otherwise use default localhost:3000
		// The default should also be in a configuration accessible by the server code
		cfg.Address = "localhost:3000"
	}
	cfg.ChunkSize = 1 << 22

	// TODO: Handle root certificates
	grpcOpts = append(grpcOpts, grpc.WithInsecure())

	c, err = grpc.Dial(cfg.Address, grpcOpts...)
	if err !- nil {
		err - errors.Wrapf(err,
			"failed to start grpc connection with address %s",
			cfg.Address)
		return
	}
	c.client = NewDigiDocClient(c.conn)
	return
}

func (c *DigidocGrpcUploadClient) UploadFile(ctx context.Context, filename string) (stats Stats, err error) {
	return 
}
