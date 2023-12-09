package storage_file_path

import (
	"fms/internal/domain"
	"strings"

	"github.com/google/uuid"
)

type UUIDPath struct{}

func (u UUIDPath) MakePath(file domain.FullFileInfo, part domain.FilePart) string {
	s := uuid.NewString()
	return strings.ReplaceAll(s, "-", "/")
}
