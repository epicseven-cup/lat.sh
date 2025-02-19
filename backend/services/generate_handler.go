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

// Handles all the request generation
func (handler GenerateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//TODO: Need to Handle bad request and make sure to send back detail error feedback so when doing cli request it can respond correctly
	data := new(types.CreateUrlRequest)
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		handler.logger.Error(err.Error())
		return
	}
	source := data.Source
	destination := data.Destination

	if source == "" {
		handler.logger.Error("Source is empty from incoming Generate request")
		http.Error(w, "Missing source", 400)
		return
	}

	if destination == "" {
		handler.logger.Error("destination is empty from incoming Generate request")
		http.Error(w, "Missing destination", 400)
		return
	}

	if handler.dbh.CheckDuplicateUrl(source) {
		handler.logger.Error("Duplicate URL on Handler")
		http.Error(w, fmt.Sprintf("Duplicate URL: %s", source), 400)
		return
	}

	handler.dbh.InsertUrl(source, destination)

	// Writing HTTP header for the response, return it as json with status
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
