package service

import (
	"api/api"
	"context"
	"fms/internal"
	"fms/internal/domain"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FileManagement struct {
	storageRepo     internal.StorageRepository
	fileRepo        internal.FileRepository
	repoSerializer  internal.RepositorySerializer
	c               FileManagementConfig
	storageRegistry *StorageRegistry
}

func NewFileManagement(
	storageRepo internal.StorageRepository,
	fileRepo internal.FileRepository,
	repoSerializer internal.RepositorySerializer,
	c FileManagementConfig,
) *FileManagement {
	return &FileManagement{
		storageRepo:     storageRepo,
		fileRepo:        fileRepo,
		repoSerializer:  repoSerializer,
		c:               c,
		storageRegistry: NewStorageRegistry(),
	}
}

func (m *FileManagement) InitializeUpload(ctx context.Context, file domain.FullFileInfo) ([]domain.FilePart, error) {
	var (
		parts []domain.FilePart
		txErr error
	)
	// нужно избежать аномалии (а именно, фантомное чтение) при работе с транзакциями, когда мы считали
	// инфу о хранилищах, распределили файлы, а к моменту записи инфы о распределении хранилища уже поменялись
	// и клиент тогда может получить ошибку нехватки места
	txErr = m.repoSerializer.Serializable(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		storages, err := m.storageRepo.GetStoragesTx(ctx, tx)
		if err != nil {
			return errors.Wrap(err, "get storages")
		}

		parts, err = m.c.PartitionStrategy.Partition(file, m.c.PartitionParts, storages)
		if err != nil {
			return errors.Wrap(err, "distribute file among storages")
		}

		for i, part := range parts {
			parts[i].Path = m.c.StorageFilePathStrategy.MakePath(file, part)
		}

		err = m.fileRepo.AddFileTx(ctx, tx, file, parts)
		if err != nil {
			return errors.Wrap(err, "save to db")
		}

		return nil
	})

	return parts, txErr
}

func (m *FileManagement) ReportUploadProgress(ctx context.Context, req domain.UploadProgress) error {
	if err := m.fileRepo.MarkPartStored(ctx, req.Filename, req.PartId); err != nil {
		return errors.Wrap(err, "mark part stored in db")
	}

	return nil
}

func (m *FileManagement) CancelUpload(ctx context.Context, filename string) error {
	parts, err := m.fileRepo.GetFileParts(ctx, filename)
	if err != nil {
		return errors.Wrap(err, "get file from db")
	}

	// Удаляем части из всех хранилищ
	var eg errgroup.Group
	for _, p := range parts {
		p := p

		eg.Go(func() error {
			storageClient, err := m.storageRegistry.GetOrAdd(ctx, p.Storage.Address)
			if err != nil {
				log.Printf("cannot get storage %s\n", p.Storage.Address)
				return errors.Wrapf(err, "get storage %s", p.Storage.Address)
			}

			if _, err := storageClient.DeleteV1(ctx, &api.DeleteV1Request{Path: p.Path}); err != nil {
				log.Printf("cannot delete from storage %s by path %s\n", p.Storage.Address, p.Path)
				return errors.Wrapf(err, "delete from storage %s by path %s", p.Storage.Address, p.Path)
			}

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	// После этого удаляем из базы (именно в таком порядке!)
	err = m.repoSerializer.Serializable(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		if err := m.fileRepo.DeleteFileTx(ctx, tx, filename); err != nil {
			return errors.Wrap(err, "delete from db")
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *FileManagement) GetFileDownloadInfo(ctx context.Context, filename string) ([]domain.FilePart, error) {
	parts, err := m.fileRepo.GetFileParts(ctx, filename)
	if err != nil {
		return nil, errors.Wrap(err, "get file from db")
	}

	for _, p := range parts {
		if !p.IsStored {
			return nil, domain.ErrFileNotCompletelyUploaded
		}
	}

	return parts, nil
}

func (m *FileManagement) AddStorage(ctx context.Context, addr string) (domain.Storage, error) {
	storageClient, err := m.storageRegistry.Add(ctx, addr)
	if err != nil {
		return domain.Storage{}, errors.Wrap(err, "cannot connect to storage")
	}

	info, err := storageClient.InfoV1(ctx, &emptypb.Empty{})
	if err != nil {
		return domain.Storage{}, errors.Wrap(err, "cannot get storage info")
	}

	st, err := m.storageRepo.AddStorage(ctx, domain.Storage{
		Address:             addr,
		SpaceAvailableBytes: info.GetSize(),
	})
	if err != nil {
		return domain.Storage{}, errors.Wrap(err, "storage repository")
	}

	return st, nil
}
