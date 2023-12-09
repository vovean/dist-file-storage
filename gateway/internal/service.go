package internal

import (
	"context"
	"gateway/internal/domain"
	"io"
)

type Service interface {
	Upload(ctx context.Context, req domain.UploadFileRequest, r io.Reader) error
	Download(ctx context.Context, req domain.DownloadFileRequest) (domain.DownloadedFile, error)
}
