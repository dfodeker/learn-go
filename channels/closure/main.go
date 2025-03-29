package main

/*
* sending a message to a closed channel will cause a panic
* A panic on the main goroutine will cause the entire program to crash, and a panic in any other goroutine will cause that goroutine to crash.
closing isnt necessary, You should close channels to indicate explicitly to a receiver that nothing else is going to come across.
At Mailio we're all about keeping track of what our systems are up to with great logging and telemetry.

The sendReports function sends out a batch of reports to our clients and reports back how many were sent across a channel. It closes the channel when it's done.

Complete the countReports function. It should:

Use an infinite for loop to read from the channel:
If the channel is closed, break out of the loop
Otherwise, keep a running total of the number of reports sent
Return the total number of reports sent




Boots

*/

func countReports(numSentCh chan int) int {

	var numSent int

	for {
		i, ok := <-numSentCh
		if !ok {
			break
		}
		numSent += i
	}

	return numSent
}

func sendReports(numBatches int, ch chan int) {
	for i := range numBatches {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
