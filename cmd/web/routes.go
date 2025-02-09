package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/saves", app.UploadSave)
	mux.HandleFunc("GET /api/saves", app.GetSave)

	return app.logRequest(mux)
}
