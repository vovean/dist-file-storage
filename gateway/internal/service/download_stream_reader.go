package service

import "api/api"

type DownloadV1Stream interface {
	Recv() (*api.DownloadV1Response, error)
}

type DownloadV1StreamReader struct {
	stream DownloadV1Stream
	buffer []byte
}

func NewGrpcStreamReader(stream DownloadV1Stream) *DownloadV1StreamReader {
	return &DownloadV1StreamReader{stream: stream}
}

func (r *DownloadV1StreamReader) Read(p []byte) (int, error) {
	if len(r.buffer) == 0 {
		msg, err := r.stream.Recv()
		if err != nil {
			return 0, err // including io.EOF
		}
		r.buffer = msg.Data
	}

	n := copy(p, r.buffer)
	r.buffer = r.buffer[n:]

	return n, nil
}
