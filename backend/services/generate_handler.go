package services

import (
	"fmt"
	"lat.sh/backend/internal/database/postgresql"
	"log/slog"
	"net/http"
)

type GenerateHandler struct {
	logger *slog.Logger
	dbh    *postgresql.PostgresDatabase
}

func (handler GenerateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		handler.logger.Error(err.Error())
		return
	}
	source := r.PostForm.Get("source")
	destination := r.PostForm.Get("destination")

	if source == "" {
		http.Error(w, "Source is empty", 400)
		handler.logger.Error("Source is empty")
		return
	}

	if destination == "" {
		http.Error(w, "destination is empty", 400)
		handler.logger.Error("destination is empty")
		return
	}

	if handler.dbh.CheckDuplicateUrl(source) {
		http.Error(w, fmt.Sprintf("Duplicate URL: %s", source), 400)
		handler.logger.Error("Duplicate URL on Handler")
		return
	}

	handler.dbh.InsertUrl(source, destination)
}

func NewGenerateHandler(logger *slog.Logger, dbh *postgresql.PostgresDatabase) GenerateHandler {
	return GenerateHandler{
		logger: logger,
		dbh:    dbh,
	}
}
