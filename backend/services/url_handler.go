package services

import (
	"encoding/json"
	"lat.sh/backend/internal/database/postgresql"
	"log/slog"
	"net/http"
)

type UrlHandler struct {
	logger *slog.Logger
	dbh    postgresql.PostgresDatabase
}

type Message struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Visitor     int    `json:"visitor"`
}

func (handler UrlHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	source := r.PathValue("url")
	tbSource, tbDestination, visitor, err := handler.dbh.SelectDestination(source)

	if err != nil {
		handler.logger.Error(err.Error())
		return
	}

	query := r.URL.Query()

	summaryFlag := false

	if query.Get("o") != "" || query.Get("overview") != "" {
		summaryFlag = true
	}

	if summaryFlag {
		data := Message{
			Source:      tbSource,
			Destination: tbDestination,
			Visitor:     visitor,
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			handler.logger.Error(err.Error())
			http.Error(w, "Do not find the url", 400)
			return
		}

	} else {
		countStatus := handler.dbh.AddVisitor(source)
		if !countStatus {
			handler.logger.Error("Fail to increase counter for visitors")
			http.Error(w, "Something went wrong", 502)
			return
		}
		http.Redirect(w, r, tbDestination, 301)
	}
}

func NewUrlHandler(logger *slog.Logger, dbh postgresql.PostgresDatabase) UrlHandler {
	return UrlHandler{
		logger,
		dbh,
	}
}
