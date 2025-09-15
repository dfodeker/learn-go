package main

import (
	"net/http"

	"github.com/dfodeker/learn-go/baby-server/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerDeleteChirp(w http.ResponseWriter, r *http.Request) {
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
	param := r.PathValue("chirpID")

	chirpID, err := uuid.Parse(param)
	if err != nil {
		respondWithError(w, http.StatusForbidden, "Invalid UUID format", err)
		return
	}
	//if the user owns the chirp allow them delete it else 403
	chirp, err := cfg.db.GetChirpByID(r.Context(), chirpID)
	if err != nil {
		//chirp not found
		respondWithError(w, http.StatusNotFound, "Chirp not found", err)
		return
	}
	//check chirp owener ship
	if chirp.UserID != userID {
		respondWithError(w, http.StatusForbidden, "User not authorized to delete this Chirp", err)
		return
	}

	err = cfg.db.DeleteChirpByID(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to delete Chirp", err)
		return
	}
	respondWithJSON(w, http.StatusNoContent, "Successfully deleted Chirp")

}
