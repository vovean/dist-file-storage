package handlers

import (
	"context"
	"errors"
	"fmt"
	"gateway/internal/domain"
	"io"
	"log"
	"net/http"
	"strconv"
)

func (h httpHandler) Download(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), h.c.Download.DownloadTimeout)
	defer cancel()

	filename := req.URL.Query().Get(filenameQueryParam)
	if filename == "" {
		http.Error(w, "empty filename", http.StatusBadRequest)
		return
	}

	r := domain.DownloadFileRequest{Filename: filename}
	resp, err := h.s.Download(ctx, r)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrFileNotFound):
			http.Error(w, domain.ErrFileNotFound.Error(), http.StatusNotFound)
		default:
			http.Error(w, ErrInternalError.Error(), http.StatusInternalServerError)
		}

		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", resp.Filename))
	w.Header().Set("Content-Length", strconv.FormatUint(resp.Size, 10))

	bytesSent, err := io.Copy(w, resp.Content)
	if err != nil {
		log.Printf("write data to client conn: %v", err)
		http.Error(w, ErrInternalError.Error(), http.StatusInternalServerError)
		return
	}

	if bytesSent != int64(resp.Size) {
		log.Printf("not all bytes written to client: %v", err)
		http.Error(w, ErrInternalError.Error(), http.StatusInternalServerError)
		return
	}
}
