package grpc

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const connTimeout = 1 * time.Second

func NewConnection(ctx context.Context, addr string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	opts = append(opts,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "grpc dial")
	}

	return conn, nil
}
