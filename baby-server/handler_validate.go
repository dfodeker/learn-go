package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type successResponse struct {
	CleanBody string `json:"cleaned_body"`
}

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error Decoding Params: %s", err)

		respondWithError(w, 500, "Something went Wrong")

		return
	}
	if len(params.Body) < 1 {
		log.Printf("Error marshalling JSON: %s", err)
		respondWithError(w, 500, "Something went Wrong")
		return
	}

	if len([]rune(params.Body)) > 140 {
		respondWithError(w, 400, "Chirp is too Long")
		return
	} else {
		fmt.Printf("Initial Value: %s\n", params.Body)

		cleanString := cleanRequestBody(params.Body)

		response := successResponse{
			CleanBody: cleanString,
		}
		fmt.Printf("val: %s\n", response)
		respondWithJson(w, 200, response)
	}

}

func cleanRequestBody(sentence string) string {
	profaneWords := []string{
		"kerfuffle",
		"sharbert",
		"fornax",
	}
	cS := sentence

	sentenceArr := strings.Split(sentence, " ")
	for i, senWord := range sentenceArr {
		s := strings.ToLower(senWord)
		for _, word := range profaneWords {
			if s == word {
				sentenceArr[i] = "****"
			}
		}
	}
	cS = strings.Join(sentenceArr, " ")

	fmt.Printf("Value: %s\n", cS)
	return cS
}
