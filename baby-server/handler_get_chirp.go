package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) getChirpByID(w http.ResponseWriter, r *http.Request) {

	param := r.PathValue("chirpID")

	if len(param) < 10 {
		respondWithError(w, http.StatusForbidden, "Please provide a valid ID")
		return
	}

	chirpID, err := uuid.Parse(param)
	if err != nil {
		respondWithError(w, http.StatusForbidden, "Invalid UUID format")
		return
	}

	chirp, err := cfg.db.GetChirpByID(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Unable to Get Chirp")
		return
	}

	resp := ChirpResponse{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		UserID:    chirp.UserID,
		Body:      chirp.Body,
	}

	respondWithJson(w, 200, resp)
}
