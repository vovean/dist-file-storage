package internal

import (
	"context"
	"dist-file-storage/internal/domain"
	"io"
)

type Service interface {
	Upload(ctx context.Context, req domain.UploadFileRequest, r io.Reader) error
	Download(ctx context.Context, req domain.DownloadFileRequest) (domain.DownloadedFile, error)
}
