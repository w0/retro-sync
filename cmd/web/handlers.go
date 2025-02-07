package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/w0/retro-sync/internal/database"
)

func (app *application) UploadSave(w http.ResponseWriter, r *http.Request) {
	save, header, err := r.FormFile("save")

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "form field save missing", err)
		return
	}

	tmp, err := os.CreateTemp("", header.Filename)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed creating temp file", err)
		return
	}

	defer tmp.Close()

	_, err = io.Copy(tmp, save)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed copying save data", err)
		return
	}

	dbSave, err := app.db.CreateSave(r.Context(), database.CreateSaveParams{
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
		Filepath:  tmp.Name(),
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to store save in database", err)
		return
	}

	respondWithJson(w, http.StatusOK, dbSave)

}

func (app *application) GetSave(w http.ResponseWriter, r *http.Request) {
	type save struct {
		Id int `json:"id"`
	}

	decoder := json.NewDecoder(r.Body)
	var s save
	err := decoder.Decode(&s)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "unexpected data in payload", err)
		return
	}

	dbSave, err := app.db.GetSave(r.Context(), int64(s.Id))

	if err != nil {
		respondWithError(w, http.StatusNotFound, "save not found", err)
		return
	}

	respondWithJson(w, http.StatusOK, dbSave)
}
