package config

import (
	"pkg/config"
	"time"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`

	Uploader struct {
		ChunkSize     config.DataSize `yaml:"chunk_size_bytes"`
		UploadTimeout time.Duration   `yaml:"upload_timeout"`
	} `yaml:"uploader"`

	Downloader struct {
		ChunkSize       config.DataSize `yaml:"chunk_size_bytes"`
		DownloadTimeout time.Duration   `yaml:"download_timeout"`
	} `yaml:"downloader"`

	FMS struct {
		Host string `yaml:"host"`
		Port uint   `yaml:"port"`
	} `yaml:"fms"`
}
