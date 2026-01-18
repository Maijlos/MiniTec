package main

import (
	"backend/src/internal/api"
	"backend/src/internal/db"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
		os.Exit(1)
	}

	database_url := os.Getenv("DATABASE_URL")
	if database_url == "" {
		slog.Error("DATABASE_URL environment variable is not set")
		os.Exit(1)
	}

	queries, err := db.New(database_url)
	if err != nil {
		slog.Error("Error connecting to database")
		os.Exit(1)
	}

	err = api.New(queries)
	if err != nil {
		slog.Error("Error starting server")
		os.Exit(1)
	}
}