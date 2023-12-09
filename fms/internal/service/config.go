package service

import "fms/internal/domain"

type FilePartitionStrategy interface {
	Partition(file domain.FullFileInfo, parts uint, storages []domain.Storage) ([]domain.FilePart, error)
}

type StorageFilePathStrategy interface {
	MakePath(file domain.FullFileInfo, part domain.FilePart) string
}

type FileManagementConfig struct {
	PartitionParts          uint
	PartitionStrategy       FilePartitionStrategy
	StorageFilePathStrategy StorageFilePathStrategy
}
