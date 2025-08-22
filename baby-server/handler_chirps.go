package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dfodeker/learn-go/baby-server/internal/database"
	"github.com/google/uuid"
)

type chirpResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	Body      string    `json:"body"`
}

func (cfg *apiConfig) CreateChirpHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body   string    `json:"body"`
		UserID uuid.UUID `json:"user_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := &parameters{}
	err := decoder.Decode(params)
	if err != nil {
		msg := fmt.Sprintf("%s", err)
		respondWithError(w, 400, "Couldn't decode parameters: "+msg)
		return
	}
	cleaned, err := validateChirp(params.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	chirpReq := database.CreateChirpParams{
		UserID: params.UserID,
		Body:   cleaned,
	}

	chirp, err := cfg.db.CreateChirp(r.Context(), chirpReq)
	if err != nil {
		respondWithError(w, 500, "Couldn't create chirp")
		return
	}
	response := chirpResponse{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		UserID:    chirp.UserID,
		Body:      chirp.Body,
	}

	respondWithJson(w, 201, response)
}

func validateChirp(body string) (string, error) {
	const maxLengthAllowed = 140
	if len(body) > maxLengthAllowed {
		return "", errors.New("chirp is too long")
	}

	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}
	cleaned := getCleanedBody(body, badWords)
	return cleaned, nil

}

func getCleanedBody(body string, badWords map[string]struct{}) string {
	words := strings.Split(body, " ")
	for i, word := range words {
		loweredWord := strings.ToLower(word)
		if _, ok := badWords[loweredWord]; ok {
			words[i] = "****"

		}
	}
	cleaned := strings.Join(words, " ")
	return cleaned
}

// func cleanRequestBody(sentence string) string {
//

// 	sentenceArr := strings.Split(sentence, " ")
// 	for i, senWord := range sentenceArr {
// 		s := strings.ToLower(senWord)
// 		for _, word := range profaneWords {
// 			if s == word {
// 				sentenceArr[i] = "****"
// 			}
// 		}
// 	}
// 	cS = strings.Join(sentenceArr, " ")

// 	fmt.Printf("Value: %s\n", cS)
// 	return cS
// }
