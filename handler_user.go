package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/thrillee/rssagg/internal/database"
)

func (apiCfg *apiConfig) handleGetUserPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetUserPosts(r.Context(), database.GetUserPostsParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error fetching user posts: %v", err))
		return
	}

	responseWithJSON(w, 200, databasePostsToPosts(posts))
}

func (apiCfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	responseWithJSON(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if params.Name == "" {
		responseWithError(w, 400, "name is required")
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:       uuid.New(),
		Created:  time.Now().UTC(),
		Modified: time.Now().UTC(),
		Name:     params.Name,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Failed creating user: %v", err))
		return
	}

	responseWithJSON(w, 201, databaseUserToUser(user))
}
