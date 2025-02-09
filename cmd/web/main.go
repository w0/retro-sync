package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/w0/retro-sync/internal/database"
	_ "modernc.org/sqlite"
)

type application struct {
	logger   *slog.Logger
	db       *database.Queries
	rootPath string
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	err := godotenv.Load(".env")
	if err != nil {
		logger.Info(err.Error())
	}

	rootPath := os.Getenv("ROOT_PATH")
	if rootPath == "" {
		logger.Error("ROOT_PATH environment variable not set.")
		os.Exit(1)
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		logger.Error("DATABASE_URL evironment variable not set.")
	}

	db, err := sql.Open("sqlite", dbUrl)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := application{
		db:       database.New(db),
		logger:   logger,
		rootPath: rootPath,
	}

	port := os.Getenv("PORT")
	if port == "" {
		app.logger.Warn("PORT environment variable unset. Default to 8080")
		port = "8080"
	}

	app.logger.Info("starting server", "PORT", port)
	http.ListenAndServe(":"+port, app.routes())
}
