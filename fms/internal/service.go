package internal

import (
	"context"
	"fms/internal/domain"
)

type FileManagementService interface {
	InitializeUpload(ctx context.Context, req domain.FullFileInfo) ([]domain.FilePart, error)
	ReportUploadProgress(ctx context.Context, req domain.UploadProgress) error
	CancelUpload(ctx context.Context, filename string) error
	GetFileDownloadInfo(ctx context.Context, filename string) ([]domain.FilePart, error)
}
