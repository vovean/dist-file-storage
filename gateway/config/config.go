package config

import (
	"pkg/config"
	"time"
)

type Config struct {
	Server struct {
		// На каком порту запускать сервер
		Port int `yaml:"port" env:"SERVER_PORT"`
	} `yaml:"server"`

	Uploader struct {
		// Размер батча (чанка) данных, которые будут вычитываться с коннекта клиента
		ChunkSize config.DataSize `yaml:"chunk_size" env:"UPLOADER_CHUNKSIZE"`
		// Максимальное время загрузки файла
		Timeout time.Duration `yaml:"timeout" env:"UPLOADER_TIMEOUT"`
	} `yaml:"uploader"`

	Downloader struct {
		// Размер батча, который будет отдаваться клиенту
		ChunkSize config.DataSize `yaml:"chunk_size" env:"DOWNLOADER_CHUNKSIZE"`
		// Максимальное время скачивания файла
		Timeout time.Duration `yaml:"timeout" env:"DOWNLOADER_TIMEOUT"`
	} `yaml:"downloader"`

	// File Management Service
	FMS struct {
		Host string `yaml:"host" env:"FMS_HOST"`
		Port uint   `yaml:"port" env:"FMS_PORT"`
	} `yaml:"fms"`
}
