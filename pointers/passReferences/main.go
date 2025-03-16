package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

func getMessageText(a *Analytics, m Message) {
	//it shoudl look at the cussess frield of messave, and based on that
	//increment the messages Total, messages succeded or messages failed
	if m.Success == true {
		a.MessagesSucceeded += 1
	} else if m.Success == false {
		a.MessagesFailed += 1
	}
	a.MessagesTotal += 1
}
