package grpc_stream

type GrpcStream[T any] interface {
	Recv() (*T, error)
}

type GrpcStreamReader[T any] struct {
	stream   GrpcStream[T]
	dataFunc func(*T) ([]byte, error)
	buffer   []byte
}

func NewGrpcStreamReader[T any](stream GrpcStream[T], dataFunc func(*T) ([]byte, error)) *GrpcStreamReader[T] {
	return &GrpcStreamReader[T]{stream: stream, dataFunc: dataFunc}
}

func (r *GrpcStreamReader[T]) Read(p []byte) (int, error) {
	if len(r.buffer) == 0 {
		msg, err := r.stream.Recv()
		if err != nil {
			return 0, err // including io.EOF
		}
		r.buffer, err = r.dataFunc(msg)
		if err != nil {
			return 0, err
		}
	}

	n := copy(p, r.buffer)
	r.buffer = r.buffer[n:]

	return n, nil
}
