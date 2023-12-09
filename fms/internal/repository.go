package internal

import (
	"context"
	"fms/internal/domain"

	"github.com/jmoiron/sqlx"
)

type FileRepository interface {
	AddFile(ctx context.Context, file domain.FullFileInfo, parts []domain.FilePart) error
	AddFileTx(ctx context.Context, tx *sqlx.Tx, file domain.FullFileInfo, parts []domain.FilePart) error
	GetFileParts(ctx context.Context, filename string) ([]domain.FilePart, error)
	DeleteFile(ctx context.Context, filename string) error
	DeleteFileTx(ctx context.Context, tx *sqlx.Tx, filename string) error
	MarkPartStored(ctx context.Context, filename string, partId int) error
}

type StorageRepository interface {
	GetStorages(ctx context.Context) ([]domain.Storage, error)
	GetStoragesTx(ctx context.Context, tx *sqlx.Tx) ([]domain.Storage, error)
	AddStorage(ctx context.Context, storage domain.Storage) (domain.Storage, error)
}

// RepositorySerializer выполняет функцию в serializable транзакции
type RepositorySerializer interface {
	Serializable(ctx context.Context, f func(ctx context.Context, tx *sqlx.Tx) error) error
}
