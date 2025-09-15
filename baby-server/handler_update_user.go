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

type userResponse struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	CreatedAt string    `json:"created_at,omitempty"`
	UpdatedAt string    `json:"updated_at,omitempty"`
}

func (cfg *apiConfig) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	bearerToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Authentication credentials are missing or invalid.", err)
		return
	}
	userID, err := auth.ValidateJWT(bearerToken, cfg.signingKey)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Authentication credentials are invalid.", err)
		return
	}
	// claims should include the authenticated user's ID

	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var params parameters
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, "bad request", err)
		return
	}

	if params.Email == "" || params.Password == "" {
		respondWithError(w, http.StatusBadRequest, "email and password are required", nil)
		return
	}
	if _, err := mail.ParseAddress(params.Email); err != nil {
		respondWithError(w, http.StatusBadRequest, "please provide a valid email", err)
		return
	}
	if _, err := mail.ParseAddress(params.Email); err != nil {
		respondWithError(w, http.StatusBadRequest, "please provide a valid email", err)
		return
	}

	hashed, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to hash password", err)
		return
	}
	//cant do this since the user is chaning theeir emails, lets see if  we can see their email tin the claim

	updateUserObject := database.UpdateUserParams{
		ID:             userID,
		Email:          params.Email,
		HashedPassword: hashed,
	}

	updated, err := cfg.db.UpdateUser(r.Context(), updateUserObject)
	if err != nil {
		// consider mapping unique-constraint/email-taken to 409
		respondWithError(w, http.StatusInternalServerError, "failed to update user", err)
		return
	}
	resp := userResponse{
		ID:        userID,
		Email:     updated.Email,
		CreatedAt: updated.CreatedAt.Format(time.RFC3339),
		UpdatedAt: updated.UpdatedAt.Format(time.RFC3339),
	}
	respondWithJSON(w, http.StatusOK, resp)

}
