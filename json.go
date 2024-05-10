package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error while json marshalling: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	errorResp := struct {
		Error string `json:"error"`
	}{
		Error: msg,
	}
	respondWithJSON(w, code, errorResp)
}
