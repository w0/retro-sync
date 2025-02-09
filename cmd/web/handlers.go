package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/w0/retro-sync/internal/database"
	"github.com/w0/retro-sync/internal/parser"
)

func (app *application) UploadSave(w http.ResponseWriter, r *http.Request) {
	systemId := r.FormValue("systemId")

	_, err := parser.ValidatePlatform(systemId)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "", err)
		return
	}

	save, header, err := r.FormFile("save")

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "form field save missing", err)
		return
	}

	filePath := app.PathBuilder("saves", systemId, header.Filename)
	if filePath == "" {
		respondWithError(w, http.StatusInternalServerError, "filePath is empty", err)
		return
	}

	if _, err := os.Stat(filePath); err != nil {
		err := os.MkdirAll(path.Dir(filePath), 0755)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "failed to create save path", err)
			return
		}
	}

	saveFile, err := os.Create(filePath)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to create file", err)
		return
	}

	defer saveFile.Close()

	_, err = io.Copy(saveFile, save)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed copying save data", err)
		return
	}

	now := time.Now()

	dbSave, err := app.db.CreateSave(r.Context(), database.CreateSaveParams{
		CreatedAt: now.Format(time.RFC3339),
		UpdatedAt: now.Format(time.RFC3339),
		SystemID:  systemId,
		Filename:  path.Base(filePath),
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
