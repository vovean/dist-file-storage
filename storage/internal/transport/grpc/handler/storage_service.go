package handler

import (
	"api/api"
	"context"
	"log"
	"pkg/grpc_stream"
	"storage/internal"
	"storage/internal/domain"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type StorageService struct {
	s internal.Service
	c StorageServiceConfig
}

func NewStorageService(s internal.Service, config StorageServiceConfig) *StorageService {
	return &StorageService{s: s, c: config}
}

func (s *StorageService) StoreV1(stream api.StorageService_StoreV1Server) error {
	metaMsg, err := stream.Recv()
	if err != nil {
		log.Println("cannot receive file metadata")
		return status.Error(codes.Internal, "internal error")
	}

	meta, ok := metaMsg.GetData().(*api.StoreV1Request_Meta)
	if !ok {
		log.Println("cannot get metadata from message")
		return status.Error(codes.Internal, "internal error")
	}

	streamReader := grpc_stream.NewGrpcStreamReader[api.StoreV1Request](
		stream,
		func(r *api.StoreV1Request) ([]byte, error) {
			data, ok := r.GetData().(*api.StoreV1Request_Content)
			if !ok {
				return nil, errors.New("message of invalid type received")
			}
			return data.Content, nil
		},
	)

	err = s.s.Store(stream.Context(), domain.StoreFileRequest{
		Content: streamReader,
		Path:    meta.Meta.GetPath(),
	})
	if err != nil {
		log.Printf("error storing file: %v\n", err)
		return status.Error(codes.Internal, "internal error")
	}

	if err := stream.SendAndClose(&emptypb.Empty{}); err != nil {
		log.Printf("error sending result and closing stream: %v", err)
		return status.Error(codes.Internal, "internal error")
	}

	return nil
}

func (s *StorageService) ServeV1(req *api.ServeV1Request, stream api.StorageService_ServeV1Server) error {
	filePart, err := s.s.Serve(stream.Context(), domain.ServeFileRequest{Path: req.GetPath()})
	if err != nil {
		if errors.Is(err, domain.ErrFileNotFound) {
			return status.Errorf(codes.NotFound, "%s not found", req.GetPath())
		}

		log.Printf("error serving file: %v\n", err)
		return status.Error(codes.Internal, "internal error")
	}

	streamWriter := grpc_stream.NewStreamWriter[*api.ServeV1Response](
		stream,
		s.c.ServeBatchSizeBytes,
		func(bytes []byte) *api.ServeV1Response {
			return &api.ServeV1Response{Data: bytes}
		},
	)

	if _, err := streamWriter.SendData(filePart); err != nil {
		log.Printf("error writing file to stream: %v\n", err)
		return status.Error(codes.Internal, "internal error")
	}

	if err := filePart.Close(); err != nil {
		// todo в продакшене подумать, хватит ли логирования или прям ошибку отдавать?
		log.Printf("cannot close file after writing to stream: %v\n", err)
	}

	return nil
}

func (s *StorageService) InfoV1(ctx context.Context, _ *emptypb.Empty) (*api.InfoV1Response, error) {
	info, err := s.s.Info(ctx)
	if err != nil {
		log.Printf("error getting storage info: %v\n", err)
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &api.InfoV1Response{Size: info.Size}, nil
}

func (s *StorageService) DeleteV1(ctx context.Context, req *api.DeleteV1Request) (*emptypb.Empty, error) {
	if err := s.s.Delete(req.GetPath()); err != nil {
		log.Printf("error deleting file: %v\n", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &emptypb.Empty{}, nil
}
