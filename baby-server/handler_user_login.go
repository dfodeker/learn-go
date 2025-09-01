package main

import (
	"encoding/json"
	"net/http"
	"net/mail"
	"time"

	"github.com/dfodeker/learn-go/baby-server/internal/auth"
)

func (cfg *apiConfig) UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email            string `json:"email"`
		Password         string `json:"password"`
		ExpiresInSeconds *int   `json:"expires_in_seconds,omitempty"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "bad request", err)
	}

	expires := 3600
	if params.ExpiresInSeconds != nil {
		expires = *params.ExpiresInSeconds
	}

	email := params.Email
	_, err = mail.ParseAddress(params.Email)
	if err != nil {
		respondWithError(w, 400, "please provide a valid email", err)
		return
	}
	user, err := cfg.db.GetUserByEmail(r.Context(), email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
	}
	//password compare
	err = auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
	}
	dur := time.Second * time.Duration(expires)
	token, err := auth.MakeJWT(user.ID, cfg.signingKey, dur)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error", err)
	}

	response := User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
		Token:     token,
	}
	respondWithJSON(w, http.StatusOK, response)

}

//st return a 401 Unauthorized response with the message "Incorrect email or password".
