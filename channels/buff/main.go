package main

/*
We want to be able to send emails in batches. A writing goroutine will write an entire batch of email messages to a buffered channel,

	and later, once the channel is full, a reading goroutine will read all of the messages from the channel and send them out to our clients.

Complete the addEmailsToQueue function. It should create a buffered channel with a buffer large enough to store all of the emails it's given.

	It should then write the emails to the channel in order, and finally return the channel.
*/
func addEmailsToQueue(emails []string) chan string {
	//ch := make(chan int, 100)
	ch := make(chan string, 100)
	for i := range emails {
		ch <- emails[i]
	}
	return ch

}

//I remembered the arrow for channels <- forgot the direction had to look it up
// since we just needed to add it to a channel we can create a loop and just add the
//each one to the channel.
