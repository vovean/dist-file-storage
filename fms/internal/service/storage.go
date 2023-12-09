package service

import (
	"api/api"
	"context"
	"log"
	"pkg/grpc"
	"sync"

	"github.com/pkg/errors"
)

var ErrStorageNotFound = errors.New("storage not connected yet")

type StorageRegistry struct {
	mu       sync.RWMutex
	storages map[string]api.StorageServiceClient
}

func NewStorageRegistry() *StorageRegistry {
	return &StorageRegistry{
		storages: make(map[string]api.StorageServiceClient),
	}
}

func (r *StorageRegistry) Get(location string) (api.StorageServiceClient, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if st, ok := r.storages[location]; ok {
		return st, nil
	}

	return nil, ErrStorageNotFound
}

func (r *StorageRegistry) Add(ctx context.Context, location string) (api.StorageServiceClient, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if st, ok := r.storages[location]; ok {
		return st, nil
	}

	conn, err := grpc.NewConnection(ctx, location)
	if err != nil {
		return nil, errors.Wrap(err, "create grpc connection")
	}

	client := api.NewStorageServiceClient(conn)
	r.storages[location] = client

	return client, nil
}

func (r *StorageRegistry) GetOrAdd(ctx context.Context, location string) (api.StorageServiceClient, error) {
	if st, err := r.Get(location); err == nil {
		return st, nil
	}

	log.Printf("opening connection to storage %s", location)
	return r.Add(ctx, location)
}
