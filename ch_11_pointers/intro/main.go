package main

import "fmt"

type Message struct {
	Recipient string
	Text      string
}

// fix the bug in the function its supposed to return a nicely formatted message to the console containing an sms's reciepient
func getMessageText(m Message) string {

	return fmt.Sprintf(`
To: %v
Message: %v
`, &m.Recipient, &m.Text)
}

// func main() {
// 	m := Message{
// 		Recipient: "Fabienne",
// 		Text:      "Zed's dead, baby. Zed's dead.",
// 	}
// 	format := getMessageText(m)
// 	fmt.Print(format)
// }
