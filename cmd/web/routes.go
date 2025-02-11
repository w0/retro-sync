package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/saves", app.UploadSave)
	mux.HandleFunc("GET /api/v1/saves", app.GetSaves)
	mux.HandleFunc("GET /api/v1/saves/{saveId}", app.GetSave)
	mux.HandleFunc("GET /api/v1/downloads/saves/{saveId}", app.DownloadSave)

	return app.logRequest(mux)
}
