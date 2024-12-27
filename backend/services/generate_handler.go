package services

import (
	"encoding/json"
	"fmt"
	"lat.sh/backend/internal/database/postgresql"
	"log/slog"
	"net/http"
)

type GenerateHandler struct {
	logger *slog.Logger
	dbh    *postgresql.PostgresDatabase
}

type Respond struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

func (handler GenerateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var respond Respond
	err := decoder.Decode(&respond)
	if err != nil {
		handler.logger.Error(err.Error())
		return
	}

	source := respond.Source
	destination := respond.Destination

	if source == "" {
		http.Error(w, "Source is empty", 400)
		handler.logger.Error("Source is empty")
		http.Error(w, "Missing source", 400)
		return
	}

	if destination == "" {
		http.Error(w, "destination is empty", 400)
		handler.logger.Error("destination is empty")
		http.Error(w, "Missing destination", 400)
		return
	}

	if handler.dbh.CheckDuplicateUrl(source) {
		http.Error(w, fmt.Sprintf("Duplicate URL: %s", source), 400)
		handler.logger.Error("Duplicate URL on Handler")
		return
	}

	handler.dbh.InsertUrl(source, destination)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text")
	_, err = w.Write([]byte("URL created"))
	if err != nil {
		handler.logger.Error(err.Error())
		return
	}
	return
}

func NewGenerateHandler(logger *slog.Logger, dbh *postgresql.PostgresDatabase) GenerateHandler {
	return GenerateHandler{
		logger: logger,
		dbh:    dbh,
	}
}
