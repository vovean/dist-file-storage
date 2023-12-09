package service

import (
	"context"
	"io"
)

// contextAwareReader оборачивает io.Reader и context.Context
type contextAwareReader struct {
	ctx    context.Context
	reader io.Reader
}

// Read реализует io.Reader для contextAwareReader
func (r *contextAwareReader) Read(p []byte) (int, error) {
	// Проверка контекста на отмену операции
	if err := r.ctx.Err(); err != nil {
		return 0, err
	}
	return r.reader.Read(p)
}
