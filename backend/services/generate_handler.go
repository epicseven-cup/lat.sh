package services

import (
	"encoding/json"
	"fmt"
	"lat.sh/backend/internal/database/postgresql"
	"lat.sh/backend/internal/types"
	"log/slog"
	"net/http"
)

type GenerateHandler struct {
	logger *slog.Logger
	dbh    *postgresql.PostgresDatabase
}

func (handler GenerateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := new(types.CreateUrlRequest)
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		handler.logger.Error(err.Error())
		return
	}
	source := data.Source
	destination := data.Destination

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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	respond := types.Respond{Message: "URL created"}
	err = json.NewEncoder(w).Encode(respond)
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
