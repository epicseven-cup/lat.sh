package main

import (
	"github.com/joho/godotenv"
	"lat.sh/backend/internal/database/postgresql"
	"lat.sh/backend/services"
	"log/slog"
	"net/http"
)

func main() {
	logger := slog.Default()
	err := godotenv.Load("../../configs/.env")

	if err != nil {
		logger.Error(err.Error())
	}

	database := postgresql.NewPostgresDatabase(*logger)
	database.Setup()

	generateHandler := services.NewGenerateHandler(logger, &database)
	http.Handle("/generate", generateHandler)

	healthHandler := services.NewHealthHandler(logger)
	http.Handle("/health", healthHandler)

	urlHandler := services.NewUrlHandler(logger, database)
	http.Handle("/redirect/{url}", urlHandler)
	logger.Error(http.ListenAndServe(":3000", nil).Error())
}
