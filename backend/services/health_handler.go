package services

import (
	"log/slog"
	"net/http"
)

type HealthHandler struct {
	logger slog.Logger
}

func (handler *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("The connection is health"))
	if err != nil {
		handler.logger.Error(err.Error())
		http.Error(w, "Health check fails", 500)
		return
	}
	handler.logger.Info("Server is healthy")
}
