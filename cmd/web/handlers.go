package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/w0/retro-sync/internal/database"
	"github.com/w0/retro-sync/internal/parser"
)

func (app *application) CreateSave(w http.ResponseWriter, r *http.Request) {

	type newSave struct {
		Filename   string `json:"fileName"`
		FileMd5    string `json:"md5"`
		PlatformId string `json:"platformId"`
	}

	decoder := json.NewDecoder(r.Body)
	var saveReq newSave
	err := decoder.Decode(&saveReq)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid json body", err)
		return
	}

	_, err = parser.ValidatePlatform(saveReq.PlatformId)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "", err)
		return
	}

	now := time.Now()

	dbSave, err := app.db.CreateSave(r.Context(), database.CreateSaveParams{
		CreatedAt: now.Format(time.RFC3339),
		UpdatedAt: now.Format(time.RFC3339),
		SystemID:  saveReq.PlatformId,
		Filename:  saveReq.Filename,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to store save in database", err)
		return
	}

	type saveRes struct {
		Id int64 `json:"id"`
	}

	respondWithJson(w, http.StatusCreated, saveRes{Id: dbSave.ID})

}

func (app *application) UploadSave(w http.ResponseWriter, r *http.Request) {
	saveId := r.PathValue("id")

	if saveId == "" {
		respondWithError(w, http.StatusBadRequest, "missing save id", nil)
		return
	}

	saveInt, err := strconv.ParseInt(saveId, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid path value", err)
		return
	}

	dbSave, err := app.db.GetSaveByID(r.Context(), saveInt)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "save not found", err)
		return
	}

	save, _, err := r.FormFile("file")

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "form field file missing", err)
		return
	}

	filePath := app.PathBuilder("saves", dbSave.SystemID, dbSave.Filename)
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

}

func (app *application) GetSave(w http.ResponseWriter, r *http.Request) {
	saveId := r.PathValue("id")

	app.logger.Debug("getting save by id", "id", saveId)

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

func (app *application) DownloadSave(w http.ResponseWriter, r *http.Request) {
	saveId := r.PathValue("id")

	app.logger.Info("serving save file", "id", saveId)

	saveInt, err := strconv.ParseInt(saveId, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid path value", err)
		return
	}

	dbSave, err := app.db.GetSaveByID(r.Context(), saveInt)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "save not found", err)
		return
	}

	localPath := app.PathBuilder("saves", dbSave.SystemID, dbSave.Filename)

	localSave, err := os.Open(localPath)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "file not found", err)
		return
	}

	defer localSave.Close()

	saveStat, err := os.Stat(localSave.Name())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error reading file", err)
		return
	}

	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", dbSave.Filename))
	w.Header().Add("Cache-Control", "no-cache")

	http.ServeContent(w, r, dbSave.Filename, saveStat.ModTime(), localSave)

}
