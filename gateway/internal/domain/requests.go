package domain

type UploadFileRequest struct {
	Filename string
	Size     int64
}

type DownloadFileRequest struct {
	Filename string
}
