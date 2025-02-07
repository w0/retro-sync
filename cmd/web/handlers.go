package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/w0/retro-sync/internal/database"
)

func (app *application) UploadSave(w http.ResponseWriter, r *http.Request) {
	save, header, err := r.FormFile("save")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tmp, err := os.CreateTemp("", header.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer tmp.Close()

	log.Println(tmp.Name())

	_, err = io.Copy(tmp, save)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dbSave, err := app.db.CreateSave(r.Context(), database.CreateSaveParams{
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
		Filepath:  tmp.Name(),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(dbSave)

	w.Write(data)

}

func (app *application) GetSave(w http.ResponseWriter, r *http.Request) {
	type save struct {
		Id int `json:"id"`
	}

	decoder := json.NewDecoder(r.Body)
	var s save
	err := decoder.Decode(&s)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbSave, err := app.db.GetSave(r.Context(), int64(s.Id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	j, _ := json.Marshal(dbSave)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
