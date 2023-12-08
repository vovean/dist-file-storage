package service

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSizedReader_Read(t *testing.T) {
	type readerParams struct {
		size        uint64
		alreadyRead uint64
		r           io.Reader
	}
	tests := []struct {
		name         string
		readerParams readerParams
		wantData     []byte
		wantN        int
		wantErr      bool
		err          error
	}{
		{
			name: "ok, size 4, len 10",
			readerParams: readerParams{
				size: 4,
				r:    bytes.NewBufferString("1234567890"),
			},
			wantN:    4,
			wantData: []byte("1234"),
			wantErr:  false,
		},
		{
			name: "ok, size 10, len 4",
			readerParams: readerParams{
				size: 10,
				r:    bytes.NewBufferString("1234"),
			},
			wantN:    4,
			wantData: []byte("1234"),
			wantErr:  false,
		},
		{
			name: "ok, size 10, len 0",
			readerParams: readerParams{
				size: 10,
				r:    bytes.NewBufferString(""),
			},
			wantN:    0,
			wantData: []byte(""),
			wantErr:  true,
			err:      io.EOF,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				r   = NewSizedReader(tt.readerParams.size, tt.readerParams.r)
				buf = make([]byte, 10)
			)

			n, err := r.Read(buf)
			if tt.wantErr {
				require.ErrorIs(t, err, tt.err)
				return
			}

			require.Equal(t, tt.wantN, n)
			require.Equal(t, tt.wantData, buf[:n])
		})
	}
}

func TestReaderPartitioner_NextPart(t *testing.T) {
	type partitionerParams struct {
		parts []uint64
		r     io.Reader
	}
	tests := []struct {
		name              string
		partitionerParams partitionerParams
		wantParts         [][]byte
	}{
		{
			name: "ok",
			partitionerParams: partitionerParams{
				parts: []uint64{1, 2, 3, 4},
				r:     bytes.NewBufferString("1234567890"),
			},
			wantParts: [][]byte{
				[]byte("1"),
				[]byte("23"),
				[]byte("456"),
				[]byte("7890"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewReaderPartitioner(tt.partitionerParams.r, tt.partitionerParams.parts)

			var (
				partsCount int
				parts      [][]byte
			)
			for {
				part, err := r.NextPart()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					t.Errorf("cannot get NextPart(): %v", err)
				}

				partsCount++

				var (
					partData []byte
					partSize int
					buf      = make([]byte, 10)
				)
				for {
					n, err := part.Read(buf)
					if err != nil {
						if errors.Is(err, io.EOF) {
							break
						}
						t.Errorf("Unexpected Read error: %v", err)
					}
					partSize += n
					partData = append(partData, buf[:n]...)
				}

				parts = append(parts, partData)
			}
			require.Equal(t, len(tt.wantParts), partsCount)
			require.Equal(t, tt.wantParts, parts)
		})
	}
}
