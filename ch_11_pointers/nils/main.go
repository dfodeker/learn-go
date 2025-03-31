package main

import (
	"strings"
)

/* update our remove profanity function
if the message is nil return early to avoid a panic
*/

func removeProfanity(message *string) {
	if message == nil {
		return
	}
	messageVal := *message

	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal
}
