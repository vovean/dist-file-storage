package service

import "io"

type SizedReader struct {
	size        uint64
	alreadyRead uint64
	r           io.Reader
}

func NewSizedReader(size uint64, r io.Reader) *SizedReader {
	return &SizedReader{size: size, r: r}
}

func (r *SizedReader) Read(buf []byte) (int, error) {
	if r.alreadyRead >= r.size {
		return 0, io.EOF
	}

	rem := r.size - r.alreadyRead
	if int(rem) < len(buf) {
		buf = buf[:rem]
	}

	n, err := r.r.Read(buf)
	r.alreadyRead += uint64(n)

	return n, err
}

type ReaderPartitioner struct {
	parts    []uint64
	r        io.Reader
	currentI int
}

func NewReaderPartitioner(r io.Reader, parts []uint64) *ReaderPartitioner {
	return &ReaderPartitioner{
		parts:    parts,
		r:        r,
		currentI: -1,
	}
}

func (r *ReaderPartitioner) NextPart() (io.Reader, error) {
	r.currentI++

	if r.currentI >= len(r.parts) {
		return nil, io.EOF
	}

	return NewSizedReader(r.parts[r.currentI], r.r), nil
}
