package config

import (
	"pkg/config"
	"time"
)

type Config struct {
	Server struct {
		Port int `yaml:"port" env:"SERVER_PORT"`
	} `yaml:"server"`

	Uploader struct {
		ChunkSize config.DataSize `yaml:"chunk_size" env:"UPLOADER_CHUNKSIZE"`
		Timeout   time.Duration   `yaml:"timeout" env:"UPLOADER_TIMEOUT"`
	} `yaml:"uploader"`

	Downloader struct {
		ChunkSize config.DataSize `yaml:"chunk_size" env:"DOWNLOADER_CHUNKSIZE"`
		Timeout   time.Duration   `yaml:"timeout" env:"DOWNLOADER_TIMEOUT"`
	} `yaml:"downloader"`

	FMS struct {
		Host string `yaml:"host" env:"FMS_HOST"`
		Port uint   `yaml:"port" env:"FMS_PORT"`
	} `yaml:"fms"`
}
