package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type errs struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	errorMsg := errs{
		Error: msg,
	}
	s, err := json.Marshal(errorMsg)
	if err != nil {
		log.Printf("Error Decoding Params: %s", err)
		http.Error(w, `{"error":"Something went wrong"}`, 500)
	}
	w.Write(s)

}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	s, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error Decoding Params: %s", err)
		http.Error(w, `{"error":"Something went wrong"}`, 500)
	}
	fmt.Printf("final Value %s\n", s)
	w.Write(s)
}
