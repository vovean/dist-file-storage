package service

import (
	"api/api"
	"context"
	goerrors "errors"
	"fmt"
	"gateway/internal/domain"
	"io"
	"log"
	"pkg/grpc_stream"
	"sort"

	"github.com/pkg/errors"
)

func (s *Service) Upload(ctx context.Context, req domain.UploadFileRequest, r io.Reader) error {
	fileParts, err := s.getUploadPlan(ctx, req.Filename, uint64(req.Size))
	if err != nil {
		return errors.Wrap(err, "get upload plan")
	}

	partsSizes := make([]uint64, 0, len(fileParts))
	for _, p := range fileParts {
		partsSizes = append(partsSizes, p.Size)
	}
	partitionedReader := NewReaderPartitioner(r, partsSizes)

	for _, p := range fileParts {
		if err := s.uploadFilePartToStorage(ctx, req.Filename, p, partitionedReader); err != nil {
			// при ошибке загрузки любой из частей отменяем загрузку в FMS
			err = goerrors.Join(err, s.cancelUpload(ctx, req.Filename))
			return errors.Wrapf(err, "send part %d of file %s to storage", p.PartId, req.Filename)
		}
	}

	return nil
}

func (s *Service) uploadFilePartToStorage(
	ctx context.Context,
	filename string,
	p domain.FilePart,
	partitionedReader *ReaderPartitioner,
) error {
	storage, err := s.storages.GetOrAdd(ctx, p.Storage)
	if err != nil {
		return errors.Wrapf(err, "get or add storage %s", p.Storage)
	}

	readerPart, err := partitionedReader.NextPart()
	if err != nil {
		return errors.Wrap(err, "get corresponding reader part")
	}

	// Создаем stream в хранилище и не забываем закрыть его потом
	stream, err := storage.StoreV1(ctx)
	if err != nil {
		return errors.Wrapf(err, "open stream to storage %s", p.Storage)
	}
	closed := false
	defer func() {
		if closed {
			return
		}
		if _, err := stream.CloseAndRecv(); err != nil {
			log.Printf("error closing stream for storage %s", p.Storage)
		}
	}()

	// Пробуем загрузить файл в хранилище.
	if err := s.sendFilePartToStorage(p, readerPart, stream); err != nil {
		return errors.Wrapf(err, "send file part (%d) to storage", p.PartId)
	}

	// Закрываем stream с хранилищем
	if _, err := stream.CloseAndRecv(); err != nil {
		return errors.Wrapf(err, "error closing stream for storage %s", p.Storage)
	}
	closed = true

	// Сообщаем FMS, что загрузили часть файла
	uploadProgress := &api.ReportUploadProgressV1Request{
		Filename: filename,
		PartId:   int32(p.PartId),
	}
	if _, err := s.fms.ReportUploadProgressV1(ctx, uploadProgress); err != nil {
		return errors.Wrapf(err, "report upload progress for part %d", p.PartId)
	}
	return nil
}

func (s *Service) cancelUpload(ctx context.Context, filename string) error {
	if _, err := s.fms.CancelUploadV1(ctx, &api.CancelUploadV1Request{Filename: filename}); err != nil {
		log.Printf("error canceling upload: %v", err)
		return err
	}
	return nil
}

func (s *Service) getUploadPlan(ctx context.Context, filename string, size uint64) ([]domain.FilePart, error) {
	uploadPlan, err := s.fms.InitFileUploadV1(ctx, &api.InitFileUploadV1Request{
		Filename: filename,
		Size:     size,
	})
	if err != nil {
		return nil, errors.Wrap(err, "initialize upload")
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

	return fileParts, nil
}

func (s *Service) sendFilePartToStorage(part domain.FilePart, partReader io.Reader, stream api.StorageService_StoreV1Client) error {
	filePartMeta := &api.StoreV1Request{Data: &api.StoreV1Request_Meta{Meta: &api.StoreRequestMetadata{Path: part.Path}}}
	if err := stream.Send(filePartMeta); err != nil {
		return errors.Wrapf(err, "send meta to storage %s", part.Storage)
	}

	streamWriter := grpc_stream.NewStreamWriter[*api.StoreV1Request](
		stream,
		s.c.UploadBatchSize,
		func(bytes []byte) *api.StoreV1Request {
			return &api.StoreV1Request{Data: &api.StoreV1Request_Content{Content: bytes}}
		},
	)

	totalSent, err := streamWriter.SendData(partReader)
	if err != nil {
		return errors.Wrapf(err, "send file part bytes to storage %s", part.Storage)
	}
	if totalSent != part.Size {
		return fmt.Errorf("sent bytes count (%d) is not equal to part size (%d)", totalSent, part.Size)
	}

	return nil
}
