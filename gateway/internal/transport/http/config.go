package http

import "time"

type DownloadEndpointConfig struct {
	DownloadTimeout time.Duration
}

type UploadEndpointConfig struct {
	UploadTimeout time.Duration
}

type Config struct {
	Port             int
	DownloadEndpoint DownloadEndpointConfig
	UploadEndpoint   UploadEndpointConfig
}
