package domain

import "io"

type StoreFileRequest struct {
	Content io.Reader
	Path    string
}

type ServeFileRequest struct {
	Path string
}

type StorageInfo struct {
	Size uint64
}
