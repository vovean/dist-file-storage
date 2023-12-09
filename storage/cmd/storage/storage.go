package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"storage/config"
	"storage/internal/service"
	"storage/internal/transport/grpc"
	"sync"
	"syscall"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	Run()
}

const defaultConfigPath = "config/storage.yml"

func Run() {
	configPath := flag.String("config", defaultConfigPath, "path to yaml config")
	flag.Parse()

	var cfg config.Config
	if err := cleanenv.ReadConfig(*configPath, &cfg); err != nil {
		log.Fatal(err)
	}

	svc, err := service.New(cfg.Storage.Root, cfg.Storage.Size.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	grpcConfig := grpc.Config{
		Port:                cfg.Server.Port,
		ServeBatchSizeBytes: cfg.API.ServeChunkSize.Bytes,
	}
	server := grpc.NewTransport(svc, grpcConfig)

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
