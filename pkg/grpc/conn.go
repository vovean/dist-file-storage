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
	ctx, cancel := context.WithTimeout(ctx, connTimeout)
	defer cancel()

	opts = append(opts,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "grpc dial %s", addr)
	}

	return conn, nil
}
