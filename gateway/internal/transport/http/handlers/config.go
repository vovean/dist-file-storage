package handlers

import "time"

type DownloadConfig struct {
	ChunkSizeBytes  int64
	DownloadTimeout time.Duration
}

type UploadConfig struct {
	UploadTimeout time.Duration
}

type Config struct {
	Upload   UploadConfig
	Download DownloadConfig
}
