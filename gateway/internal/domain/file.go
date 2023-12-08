package domain

import "io"

type FilePart struct {
	PartId  int    // порядковый номер части файла среди всех его частей
	Storage string // адрес сервера-хранилища
	Size    uint64
	Path    string // путь на сервере
}

type FileParts []FilePart

func (ps FileParts) TotalSize() uint64 {
	var s uint64
	for _, p := range ps {
		s += p.Size
	}
	return s
}

type DownloadedFile struct {
	Size     uint64
	Filename string
	Content  io.Reader
}
