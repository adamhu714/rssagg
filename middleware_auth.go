package main

import (
	"fmt"
	"net/http"

	"github.com/adamhu714/rssagg/internal/auth"
	"github.com/adamhu714/rssagg/internal/database"
)



type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Unauthorised request: %s", err.Error()))
		return
	}
	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apikey)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get user: %s", err.Error()))
		return
	}
	handler(w, r, user)
	}
}
