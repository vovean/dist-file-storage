package service

import (
	"api/api"
	"context"
	"pkg/grpc"

	"github.com/pkg/errors"
)

type Config struct {
	FMSAddr           string
	UploadBatchSize   uint64
	DownloadBatchSize uint64
}

type Service struct {
	fms      api.FileManagementServiceClient
	storages *StorageRegistry
	c        Config
}

func New(ctx context.Context, c Config) (*Service, error) {
	fmsConn, err := grpc.NewConnection(ctx, c.FMSAddr)
	if err != nil {
		return nil, errors.Wrap(err, "create grpc connection")
	}

	fms := api.NewFileManagementServiceClient(fmsConn)

	return &Service{
		fms:      fms,
		storages: NewStorageRegistry(),
		c:        c,
	}, nil
}
