package main

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	// ?
	//make sure required fields have non zero values
	if mToSend.sender.name == "" {
		return false
	} else if mToSend.sender.number == 0 {
		return false
	} else if mToSend.recipient.name == "" {
		return false
	} else if mToSend.recipient.number == 0 {
		return false
	}
	//return true only if the sender and reciepend fields
	//each constain a name and a number.
	//if a default 0 value is used return false
	return true
}
