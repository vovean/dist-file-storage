package service

import (
	"api/api"
	"context"
	"dist-file-storage/internal/domain"
	"io"
	"sort"
	"sync"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func (s *Service) Download(ctx context.Context, req domain.DownloadFileRequest) (domain.DownloadedFile, error) {
	downloadPlan, err := s.fms.GetFileDownloadInfoV1(ctx, &api.GetFileDownloadInfoV1Request{Filename: req.Filename})
	if err != nil {
		return domain.DownloadedFile{}, errors.Wrap(err, "initialize download")
	}

	fileParts := make([]domain.FilePart, 0, len(downloadPlan.GetFileParts()))
	for _, fp := range downloadPlan.GetFileParts() {
		fileParts = append(fileParts, domain.FilePart{
			PartId:  int(fp.GetPartId()),
			Storage: fp.GetStorage(),
			Size:    fp.GetSize(),
			Path:    fp.GetPath(),
		})
	}
	sort.Slice(fileParts, func(i, j int) bool {
		return fileParts[i].PartId < fileParts[j].PartId
	})

	eg, ctx := errgroup.WithContext(ctx)

	var (
		partDataReaders = make(map[int]io.Reader, len(fileParts))
		mu              sync.Mutex
	)
	for _, fp := range fileParts {
		fp := fp
		eg.Go(func() error {
			storage, err := s.storages.GetOrAdd(ctx, fp.Storage)
			if err != nil {
				return errors.Wrapf(err, "get or add storage %s", fp.Storage)
			}

			stream, err := storage.DownloadV1(ctx, &api.DownloadV1Request{Path: fp.Path})
			if err != nil {
				return errors.Wrapf(err, "open download stream %s", fp.Storage)
			}

			streamReader := NewGrpcStreamReader(stream)

			mu.Lock()
			partDataReaders[fp.PartId] = streamReader
			mu.Unlock()

			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return domain.DownloadedFile{}, errors.Wrap(err, "parallel download error")
	}

	inOrderReaders := make([]io.Reader, 0, len(fileParts))
	for _, fp := range fileParts {
		inOrderReaders[fp.PartId] = partDataReaders[fp.PartId]
	}
	fileReader := io.MultiReader(inOrderReaders...)

	return domain.DownloadedFile{
		Size:    domain.FileParts(fileParts).TotalSize(),
		Content: fileReader,
	}, nil
}
