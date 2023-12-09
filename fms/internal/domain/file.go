package domain

type FullFileInfo struct {
	Filename string
	Size     uint64
}

type FilePart struct {
	PartId   int
	Storage  Storage
	Size     uint64
	Path     string
	IsStored bool
}

type UploadProgress struct {
	Filename string
	PartId   int
}
