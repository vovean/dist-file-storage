package handler

import (
	"api/api"
	"context"
	"errors"
	"fms/internal"
	"fms/internal/domain"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FileManagementService struct {
	s internal.FileManagementService
}

func NewFileManagementService(s internal.FileManagementService) *FileManagementService {
	return &FileManagementService{s: s}
}

func (fms *FileManagementService) InitFileUploadV1(ctx context.Context, req *api.InitFileUploadV1Request) (*api.InitFileUploadV1Response, error) {
	parts, err := fms.s.InitializeUpload(ctx, domain.FullFileInfo{
		Filename: req.GetFilename(),
		Size:     req.GetSize(),
	})
	if err != nil {
		log.Printf("cannot initialize upload: %v", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	resp := &api.InitFileUploadV1Response{FileParts: make([]*api.FilePart, 0, len(parts))}
	for _, p := range parts {
		resp.FileParts = append(resp.FileParts, &api.FilePart{
			PartId:  int32(p.PartId),
			Storage: p.Storage.Address,
			Size:    p.Size,
			Path:    p.Path,
		})
	}

	return resp, nil
}

func (fms *FileManagementService) ReportUploadProgressV1(ctx context.Context, req *api.ReportUploadProgressV1Request) (*emptypb.Empty, error) {
	err := fms.s.ReportUploadProgress(ctx, domain.UploadProgress{
		Filename: req.GetFilename(),
		PartId:   int(req.GetPartId()),
	})
	if err != nil {
		log.Printf("cannot report upload progress: %v", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &emptypb.Empty{}, nil
}

func (fms *FileManagementService) GetFileDownloadInfoV1(ctx context.Context, req *api.GetFileDownloadInfoV1Request) (*api.GetFileDownloadInfoV1Response, error) {
	parts, err := fms.s.GetFileDownloadInfo(ctx, req.GetFilename())
	if err != nil {
		log.Printf("cannot get file info: %v", err)
		if errors.Is(err, domain.ErrFileNotFound) {
			return nil, status.Error(codes.NotFound, "file not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	resp := &api.GetFileDownloadInfoV1Response{FileParts: make([]*api.FilePart, 0)}
	for _, p := range parts {
		resp.FileParts = append(resp.FileParts, &api.FilePart{
			PartId:  int32(p.PartId),
			Storage: p.Storage.Address,
			Size:    p.Size,
			Path:    p.Path,
		})
	}

	return resp, nil
}

func (fms *FileManagementService) CancelUploadV1(ctx context.Context, req *api.CancelUploadV1Request) (*emptypb.Empty, error) {
	// todo по-хорошему, операцию отмены загрузки я б кидал через кафку (быстрее для клиента, надежнее для нас),
	// 	    но для скорости разработки в рамках тестового сделаем через grpc
	err := fms.s.CancelUpload(ctx, req.GetFilename())
	if err != nil {
		log.Printf("cannot cancel upload: %v", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &emptypb.Empty{}, nil
}
