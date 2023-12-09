package grpc_stream

import (
	"io"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

type Stream[T proto.Message] interface {
	Send(T) error
}

// StreamWriter - шаблонизированная структура для отправки данных в gRPC stream.
type StreamWriter[T proto.Message] struct {
	stream      Stream[T]
	batchSize   uint64
	messageFunc func([]byte) T
}

func NewStreamWriter[T proto.Message](stream Stream[T], batchSize uint64, messageFunc func([]byte) T) *StreamWriter[T] {
	return &StreamWriter[T]{
		stream:      stream,
		batchSize:   batchSize,
		messageFunc: messageFunc,
	}
}

// SendData читает данные из io.Reader и отправляет их в gRPC stream.
func (s *StreamWriter[T]) SendData(r io.Reader) (uint64, error) {
	var (
		buffer    = make([]byte, s.batchSize)
		totalSent uint64
	)
	for {
		n, err := r.Read(buffer)
		if n > 0 {
			message := s.messageFunc(buffer[:n])
			if err := s.stream.Send(message); err != nil {
				return 0, err
			}
			totalSent += uint64(n)
		}
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return 0, err
		}
	}

	return totalSent, nil
}
