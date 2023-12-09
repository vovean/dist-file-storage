package storage_file_path

import (
	"fms/internal/domain"
	"strings"

	"github.com/google/uuid"
)

type UUIDPath struct{}

func (u UUIDPath) MakePath(_ domain.FullFileInfo, _ domain.FilePart) string {
	s := uuid.NewString()
	return strings.ReplaceAll(s, "-", "/")
}
