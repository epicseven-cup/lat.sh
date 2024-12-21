package main

import (
	"lat.sh/backend/internal/database/postgresql"
	"log/slog"
	"time"
)

const Address = ":8080"

const ReadTimeout = 10 * time.Second

const WriteTimeout = 10 * time.Second

func main() {
	logger := slog.Default()
	database := postgresql.NewPostgresDatabase(*logger)
	database.Setup()
}
