package service

import (
	"api/api"
	"context"
	"fmt"
	"gateway/internal/domain"
	"io"
	"log"
	"sort"

	"github.com/pkg/errors"
)

func (s *Service) Upload(ctx context.Context, req domain.UploadFileRequest, r io.Reader) error {
	uploadPlan, err := s.fms.InitFileUploadV1(ctx, &api.InitFileUploadV1Request{
		Filename: req.Filename,
		Size:     uint64(req.Size),
	})
	if err != nil {
		return errors.Wrap(err, "initialize upload")
	}

	fileParts := make([]domain.FilePart, 0, len(uploadPlan.GetFileParts()))
	for _, p := range uploadPlan.GetFileParts() {
		fileParts = append(fileParts, domain.FilePart{
			PartId:  int(p.GetPartId()),
			Storage: p.GetStorage(),
			Size:    p.GetSize(),
			Path:    p.GetPath(),
		})
	}
	sort.Slice(fileParts, func(i, j int) bool {
		return fileParts[i].PartId < fileParts[j].PartId
	})

	partsSizes := make([]uint64, 0, len(fileParts))
	for _, p := range fileParts {
		partsSizes = append(partsSizes, p.Size)
	}
	partitionedReader := NewReaderPartitioner(r, partsSizes)

	for _, p := range fileParts {
		p := p // чтобы использовать в defer ниже

		storage, err := s.storages.GetOrAdd(ctx, p.Storage)
		if err != nil {
			return errors.Wrapf(err, "get or add storage %s", p.Storage)
		}

		readerPart, err := partitionedReader.NextPart()
		if err != nil {
			return errors.Wrap(err, "get corresponding reader part")
		}

		stream, err := storage.StoreV1(ctx)
		if err != nil {
			return errors.Wrapf(err, "open stream to storage %s", p.Storage)
		}
		closed := false
		defer func() { // Goland ругвется, но здесь правда нужен defer в цикле, чтоб закрыть все стримы
			if closed {
				return
			}
			if _, err := stream.CloseAndRecv(); err != nil {
				log.Printf("error closing stream for storage %s", p.Storage)
			}
		}()

		if err := s.sendFilePartToStorage(p, readerPart, stream); err != nil {
			return errors.Wrapf(err, "send file part (%d) to storage", p.PartId)
		}

		if _, err := stream.CloseAndRecv(); err != nil {
			return errors.Wrapf(err, "error closing stream for storage %s", p.Storage)
		}
		closed = true
	}

	return nil
}

func (s *Service) sendFilePartToStorage(part domain.FilePart, partReader io.Reader, stream api.StorageService_StoreV1Client) error {
	filePartMeta := &api.StoreV1Request{Data: &api.StoreV1Request_Meta{Meta: &api.StoreRequestMetadata{Path: part.Path}}}
	if err := stream.Send(filePartMeta); err != nil {
		return errors.Wrapf(err, "send meta to storage %s", part.Storage)
	}

	var (
		buf       = make([]byte, s.c.UploadBatchSize)
		totalSent uint64
	)
	for {
		n, err := partReader.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return errors.Wrapf(err, "read from reader part %d", part.PartId)
		}

		if err := stream.Send(&api.StoreV1Request{Data: &api.StoreV1Request_Content{Content: buf[:n]}}); err != nil {
			return errors.Wrapf(err, "send file part bytes to storage %s", part.Storage)
		}

		totalSent += uint64(n)
	}

	if totalSent != part.Size {
		return fmt.Errorf("read bytes (%d) is not equal to part size (%d)", totalSent, part.Size)
	}

	return nil
}
