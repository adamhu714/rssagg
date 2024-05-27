package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/adamhu714/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerPostsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	limitStr := r.URL.Query().Get("limit")
	limit := 5
	if limitStr != "" {
		limitConv, err := strconv.Atoi(limitStr)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Sprintf("error converting limit query parameter to int: %s", err.Error()))
			return
		}
		limit = limitConv
	}
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(limit),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error getting posts from database: %s", err.Error()))
		return
	}
	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
