package main

import (
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
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
	saveId := r.PathValue("saveId")

	app.logger.Debug("getting save by id", "saveId", saveId)

	saveInt, err := strconv.ParseInt(saveId, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid save id", err)
		return
	}

	dbSave, err := app.db.GetSaveByID(r.Context(), saveInt)

	if err != nil {
		respondWithError(w, http.StatusNotFound, "save not found", err)
		return
	}

	respondWithJson(w, http.StatusOK, dbSave)
}

func (app *application) GetSaves(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()

	page := queryValues.Get("page")
	if page == "" {
		page = "0"
	}

	pageSize := queryValues.Get("page-size")
	if pageSize == "" {
		pageSize = "100"
	}

	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid page", err)
		return
	}

	pageSizeInt, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid page-size", err)
		return
	}

	dbSaves, err := app.db.GetSaves(r.Context(), database.GetSavesParams{
		Offset: pageInt,
		Limit:  pageSizeInt,
	})

	respondWithJson(w, http.StatusOK, dbSaves)

}
