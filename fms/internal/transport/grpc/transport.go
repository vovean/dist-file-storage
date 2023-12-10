package grpc

import (
	"api/api"
	"context"
	"fms/internal"
	"fms/internal/transport/grpc/handler"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const serverShutdownTimeout = 30 * time.Second // todo в конфиг

type Transport struct {
	server *grpc.Server
	c      Config
}

func NewTransport(
	s internal.FileManagementService,
	admS internal.FileManagementAdminService,
	c Config,
) *Transport {
	server := grpc.NewServer()

	reflection.Register(server)

	fmsHandler := handler.NewFileManagementService(s)
	fmasHandler := handler.NewFileManagementAdminService(admS)
	api.RegisterFileManagementServiceServer(server, fmsHandler)
	api.RegisterFileManagementAdminServiceServer(server, fmasHandler)

	return &Transport{server: server, c: c}
}

func (t *Transport) Run(ctx context.Context) error {
	transportCtx, transportCancel := context.WithCancel(ctx)
	go func() {
		defer transportCancel()

		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", t.c.Port))
		if err != nil {
			log.Printf("failed to listen: %v", err)
			return
		}

		log.Printf("Server started at %d\n", t.c.Port)
		if err := t.server.Serve(lis); err != nil {
			if errors.Is(err, grpc.ErrServerStopped) {
				log.Println("Server stopped")
				return
			}
			log.Printf("Server error: %v\n", err)
		}
	}()

	select {
	case <-ctx.Done():
	case <-transportCtx.Done():
	}

	ctx, cancel := context.WithTimeout(context.Background(), serverShutdownTimeout)
	defer cancel()

	stopDone := make(chan struct{})
	go func() {
		t.server.GracefulStop()
		close(stopDone)
	}()

	select {
	case <-ctx.Done():
		return errors.New("server shutdown incorrectly by timeout")
	case <-stopDone:
		return errors.New("server shutdown correctly")
	}
}
