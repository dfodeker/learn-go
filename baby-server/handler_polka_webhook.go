package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dfodeker/learn-go/baby-server/internal/auth"
	"github.com/google/uuid"
)

// {
//   "event": "user.upgraded",
//   "data": {
//     "user_id": "3311741c-680c-4546-99f3-fc9efac2036c"
//   }
// }

type UpgradeEvent struct {
	Event string `json:"event"`
	Data  eventData
}
type eventData struct {
	UserID string `json:"user_id"`
}

const upgradeString = "user.upgraded"

func (cfg *apiConfig) polkaWebhookHandler(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Authentication credentials are missing or invalid v.", err)
		return
	}

	if apiKey != cfg.polkaApiKey {
		log.Printf("Polka Key:%s", cfg.polkaApiKey)
		respondWithError(w, http.StatusUnauthorized, "Authentication credentials are missing or invalid.", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	event := UpgradeEvent{}
	err = decoder.Decode(&event)
	if err != nil {
		log.Printf("Error Decoding Params: %s", err)
		respondWithError(w, 400, "Please provide a valid request body", err)
		return
	}
	if event.Event != upgradeString {
		respondWithError(w, http.StatusNoContent, "Please provide a valid request body", err)
	}
	if event.Data.UserID == "" {
		respondWithError(w, http.StatusNotFound, "Invalid UUID format", err)
		return

	}
	id, err := uuid.Parse(event.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Invalid UUID format", err)
		return
	}
	if event.Event == upgradeString {
		//mark as chirpy red
		_, err := cfg.db.UpgradeUserToChirpyRed(r.Context(), id)
		if err != nil {
			log.Printf("Error Upgrading User %s: %s", event.Data.UserID, err)
			respondWithError(w, http.StatusForbidden, "Could not Update user ", err)
			return
		}
		respondWithJSON(w, http.StatusNoContent, "")
		return

	}
}
