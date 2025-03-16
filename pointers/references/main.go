package main

import (
	"strings"
)

/*
>
*/

func removeProfanity(message *string) {
	badWords := map[string]string{
		"fubb":  "****",
		"shiz":  "****",
		"witch": "*****",
	}

	for key, word := range badWords {
		*message = strings.ReplaceAll(*message, key, word)
	}

}
