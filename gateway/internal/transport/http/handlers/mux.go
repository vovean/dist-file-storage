package handlers

import (
	"dist-file-storage/internal"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateMux(s internal.Service, c Config) http.Handler {
	handler := httpHandler{
		s: s,
		c: c,
	}

	m := mux.NewRouter()
	m.Path("/api/v1/file").Methods("POST").HandlerFunc(handler.Upload)
	m.Path("/api/v1/file").Methods("GET").HandlerFunc(handler.Download)

	return m
}

type httpHandler struct {
	s internal.Service
	c Config
}
