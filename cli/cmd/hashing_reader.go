package cmd

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
	"io"
)

type hashingReader struct {
	r      io.Reader
	hasher hash.Hash
}

func newHashingReader(reader io.Reader) *hashingReader {
	return &hashingReader{
		r:      reader,
		hasher: md5.New(),
	}
}

func (hr *hashingReader) Read(p []byte) (int, error) {
	n, err := hr.r.Read(p)
	if n > 0 {
		_, err := hr.hasher.Write(p[:n])
		if err != nil {
			return 0, err
		}
	}
	return n, err
}

// Hash возвращает текущий хэш в виде строки
func (hr *hashingReader) Hash() string {
	return hex.EncodeToString(hr.hasher.Sum(nil))
}
