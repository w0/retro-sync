package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/saves", app.CreateSave)
	mux.HandleFunc("POST /api/v1/saves/{id}/upload", app.UploadSave)
	mux.HandleFunc("GET /api/v1/saves", app.GetSaves)
	mux.HandleFunc("GET /api/v1/saves/{id}", app.GetSave)
	mux.HandleFunc("GET /api/v1/downloads/saves/{id}", app.DownloadSave)
	mux.HandleFunc("POST /api/v1/games", app.CreateGame)

	return app.logRequest(mux)
}
