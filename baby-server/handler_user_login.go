package main

import (
	"encoding/json"
	"net/http"
	"net/mail"
	"time"

	"github.com/dfodeker/learn-go/baby-server/internal/auth"
	"github.com/dfodeker/learn-go/baby-server/internal/database"
	"github.com/google/uuid"
)

type RefreshTokenParams struct {
	Token     string
	UserID    uuid.UUID
	ExpiresAt time.Time
}

func (cfg *apiConfig) UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "bad request", err)
	}

	expires := 3600

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
	refreshTime := time.Now().Add(60 * 24 * time.Hour) // 60 days
	//token, user_id, expires_at
	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error", err)
		return
	}

	dur := time.Second * time.Duration(expires)
	token, err := auth.MakeJWT(user.ID, cfg.signingKey, dur)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error", err)
		return
	}
	refresh, err := cfg.db.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		Token:     refreshToken,
		UserID:    user.ID,
		ExpiresAt: refreshTime,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to generate refresh toekn", err)
		return
	}

	response := User{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Email:        user.Email,
		Token:        token,
		RefreshToken: refresh.Token,
	}
	respondWithJSON(w, http.StatusOK, response)

}

//st return a 401 Unauthorized response with the message "Incorrect email or password".
