package handlers

import (
	"context"
	"errors"
	"gateway/internal/domain"
	"log"
	"net/http"
)

const filenameQueryParam = "filename"

func (h httpHandler) Upload(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), h.c.Upload.UploadTimeout)
	defer cancel()

	contentLength := req.ContentLength
	if contentLength <= 0 {
		http.Error(w, "Content-Length not set", http.StatusLengthRequired)
		return
	}

	filename := req.URL.Query().Get(filenameQueryParam)
	if filename == "" {
		http.Error(w, "empty filename", http.StatusBadRequest)
		return
	}

	r := domain.UploadFileRequest{
		Filename: filename,
		Size:     contentLength,
	}
	err := h.s.Upload(ctx, r, req.Body)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrFileAlreadyExists):
			http.Error(w, domain.ErrFileAlreadyExists.Error(), http.StatusConflict)
		default:
			log.Printf("upload file: %v", err)
			http.Error(w, "", http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(http.StatusCreated)
}
