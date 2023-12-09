package main

import (
	"context"
	"flag"
	"fmt"
	"gateway/config"
	"gateway/internal/service"
	"gateway/internal/transport/http"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	Run()
}

const defaultConfigPath = "config/gateway.yml"

func Run() {
	configPath := flag.String("config", defaultConfigPath, "path to yaml config")
	flag.Parse()

	var cfg config.Config
	if err := cleanenv.ReadConfig(*configPath, &cfg); err != nil {
		log.Fatal(err)
	}

	serviceConfig := service.Config{
		FMSAddr:           fmt.Sprintf("%s:%d", cfg.FMS.Host, cfg.FMS.Port),
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := server.Run(ctx); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	<-c
	cancel()

	wg.Wait()
	log.Println("all goroutines stopped, shutting down...")
}
