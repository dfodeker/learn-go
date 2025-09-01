package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) getChirpByID(w http.ResponseWriter, r *http.Request) {

	param := r.PathValue("chirpID")

	chirpID, err := uuid.Parse(param)
	if err != nil {
		respondWithError(w, http.StatusForbidden, "Invalid UUID format", err)
		return
	}

	chirp, err := cfg.db.GetChirpByID(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Unable to Get Chirp", err)
		return
	}

	resp := ChirpResponse{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		UserID:    chirp.UserID,
		Body:      chirp.Body,
	}

	respondWithJSON(w, 200, resp)
}
