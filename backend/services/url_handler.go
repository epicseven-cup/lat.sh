package services

import (
	"log/slog"
	"net/http"
)

type UrlHandler struct {
	logger slog.Logger
}

func (handler *UrlHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
