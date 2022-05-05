package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"unicode/utf8"
)

type hash struct {
	Hash string `json:hash`
}

func (a *App) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) respondWithError(w http.ResponseWriter, code int, message string) {
	a.respondWithJSON(w, code, map[string]string{"error": message})

	a.Logger.Printf("App error: code %d, message %s", code, message)
}

func (a *App) sendCommitHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash_len := utf8.RuneCountInString(vars["hash"])
	if hash_len != 40 {
		a.respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Expected hash of 40 characters, got %v", hash_len))
	} else {
		a.respondWithJSON(w, http.StatusOK, vars)
	}

}
