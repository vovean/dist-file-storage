package handler

import (
	"api/api"
	"context"
	"fms/internal"
	"fms/internal/domain"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FileManagementAdminService struct {
	s internal.FileManagementAdminService
}

func NewFileManagementAdminService(s internal.FileManagementAdminService) *FileManagementAdminService {
	return &FileManagementAdminService{s: s}
}

func (s *FileManagementAdminService) AddStorageV1(ctx context.Context, req *api.AddStorageV1Request) (*api.AddStorageV1Response, error) {
	st, err := s.s.AddStorage(ctx, domain.Storage{
		Address:             req.GetAddr(),
		SpaceAvailableBytes: req.GetSpaceAvailableBytes(),
	})
	if err != nil {
		log.Printf("error adding storage: %v\n", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &api.AddStorageV1Response{Id: int32(st.Id)}, nil
}
