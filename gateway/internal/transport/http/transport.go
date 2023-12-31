package http

import (
	"context"
	"fmt"
	"gateway/internal"
	"gateway/internal/transport/http/handlers"
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const serverShutdownTimeout = 30 * time.Second // todo в конфиг
const readHeaderTimeout = 1 * time.Second      // todo в конфиг

type Transport struct {
	server *http.Server
}

func NewTransport(s internal.Service, c Config) *Transport {
	mux := handlers.CreateMux(s, handlers.Config{
		Upload: handlers.UploadConfig{
			UploadTimeout: c.UploadEndpoint.UploadTimeout,
		},
		Download: handlers.DownloadConfig{
			DownloadTimeout: c.DownloadEndpoint.DownloadTimeout,
		},
	})

	return &Transport{
		server: &http.Server{
			Addr:              fmt.Sprintf(":%d", c.Port),
			ReadHeaderTimeout: readHeaderTimeout,
			Handler:           mux,
		},
	}
}

func (t *Transport) Run(ctx context.Context) error {
	transportCtx, transportCancel := context.WithCancel(ctx)
	go func() {
		defer transportCancel()

		log.Printf("Server started at %s\n", t.server.Addr)
		if err := t.server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
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

	err := t.server.Shutdown(ctx)

	return errors.Wrap(err, "server shutdown")
}
