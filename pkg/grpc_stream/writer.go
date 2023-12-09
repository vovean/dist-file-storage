package grpc_stream

import (
	"io"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// StreamWriter - шаблонизированная структура для отправки данных в gRPC stream.
type StreamWriter[T any] struct {
	stream      grpc.ServerStream
	batchSize   uint64
	messageFunc func([]byte) T
}

func NewStreamWriter[T any](stream grpc.ServerStream, batchSize uint64, messageFunc func([]byte) T) *StreamWriter[T] {
	return &StreamWriter[T]{
		stream:      stream,
		batchSize:   batchSize,
		messageFunc: messageFunc,
	}
}

// SendData читает данные из io.Reader и отправляет их в gRPC stream.
func (s *StreamWriter[T]) SendData(r io.Reader) error {
	buffer := make([]byte, s.batchSize)

	for {
		n, err := r.Read(buffer)
		if n > 0 {
			message := s.messageFunc(buffer[:n])
			if err := s.stream.SendMsg(message); err != nil {
				return err
			}
		}
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}
	}

	return nil
}
