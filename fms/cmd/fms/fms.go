package main

import (
	"context"
	"flag"
	"fms/config"
	"fms/internal/repository"
	"fms/internal/service"
	"fms/internal/service/file_partitioning"
	"fms/internal/service/storage_file_path"
	"fms/internal/transport/grpc"
	"log"
	"os"
	"os/signal"
	postgresql "pkg/postgres"
	"sync"
	"syscall"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	Run()
}

const defaultConfigPath = "config/fms.yml"

func Run() {
	configPath := flag.String("config", defaultConfigPath, "path to yaml config")
	flag.Parse()

	var cfg config.Config
	if err := cleanenv.ReadConfig(*configPath, &cfg); err != nil {
		log.Fatal(err)
	}

	// todo в рамках тестового задания не стал делать полноценный DI слой, а вот прям тут все делаю
	//      в проде, конечно, DI

	// База данных
	dbConfig := postgresql.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Dbname:   cfg.DB.Dbname,
		SSLMode:  cfg.DB.SSLMode,
	}
	db, err := dbConfig.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Сервис
	fmsRepository := repository.NewFMSRepository(db)
	svc := service.NewFileManagement(
		fmsRepository,
		fmsRepository,
		fmsRepository,
		service.FileManagementConfig{
			PartitionParts:          cfg.FileParts,
			PartitionStrategy:       file_partitioning.ByFreeSpaceStrategy{},
			StorageFilePathStrategy: storage_file_path.UUIDPath{},
		})

	// Сервер
	grpcConfig := grpc.Config{
		Port: cfg.Server.Port,
	}
	server := grpc.NewTransport(
		svc,
		svc,
		grpcConfig,
	)

	// Запускаем
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

	// Graceful shutdown
	c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-c
	cancel()

	wg.Wait()
	log.Println("all goroutines stopped, shutting down...")
}
