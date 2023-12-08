package main

import (
	"context"
	"dist-file-storage/config"
	"dist-file-storage/internal/service"
	"dist-file-storage/internal/transport/http"
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	Run()
}

const configPath = "config/gateway.yml"

func Run() {
	var cfg config.Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal(err)
	}

	serviceConfig := service.Config{
		FMSAddr:           fmt.Sprintf("%s:%s", cfg.FMS.Host, cfg.FMS.Port),
		UploadBatchSize:   cfg.Uploader.ChunkSize.Bytes,
		DownloadBatchSize: cfg.Downloader.ChunkSize.Bytes,
	}
	s, err := service.New(context.Background(), serviceConfig)
	if err != nil {
		log.Fatal(err)
	}

	serverConfig := http.Config{
		Port: cfg.Server.Port,
		DownloadEndpoint: http.DownloadEndpointConfig{
			DownloadTimeout: cfg.Downloader.DownloadTimeout,
		},
		UploadEndpoint: http.UploadEndpointConfig{
			UploadTimeout: cfg.Uploader.UploadTimeout,
		},
	}
	server := http.NewTransport(s, serverConfig)

	if err := server.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
