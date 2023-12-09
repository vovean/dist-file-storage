package internal

import (
	"context"
	"io"
	"storage/internal/domain"
)

type Service interface {
	Store(ctx context.Context, req domain.StoreFileRequest) error
	Serve(ctx context.Context, req domain.ServeFileRequest) (io.ReadCloser, error)
	Info(ctx context.Context) (domain.StorageInfo, error)
	Delete(path string) error
}
