package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/w0/retro-sync/internal/database"
	_ "modernc.org/sqlite"
)

type application struct {
	db *database.Queries
}

func main() {
	db, err := sql.Open("sqlite", "./retrosync.db?_pragma=foreign_keys(1)")

	if err != nil {
		log.Fatalf("failed to open db: %s", err)
	}

	defer db.Close()

	app := application{
		db: database.New(db),
	}

	log.Println("starting server on :4000")
	http.ListenAndServe(":4000", app.routes())
}
