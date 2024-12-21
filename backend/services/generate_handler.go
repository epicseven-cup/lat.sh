package services

import (
	"log/slog"
	"net/http"
)

type GenerateHandler struct {
	logger slog.Logger
}

func (handler *GenerateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
