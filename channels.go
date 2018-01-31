package main

import (
	"fmt"
)

func process(messages chan string, quit chan struct{}) {
	// Loop through your messages
	for m := range messages {
		// Check to see if they sent a blank string, if so , quit
		if m == "" {
			break
		}

		// print the message
		fmt.Println(m)
	}
	close(quit)
}

func main() {
	// declare the messages channel of type string and capacity of 5
	messages := make(chan string, 5)

	// declare a signal channel
	quit := make(chan struct{})

	// launch process in a goroutine
	go process(messages, quit)

	// declare 5 fruits in a []string
	fruits := []string{"Apple", "Kiwi", "Banana", "Orange", "Berry"}
	for _, f := range fruits {
		// loop through the fruits and send them to the messages channel
		messages <- f
	}

	// write an empty string to the messages channel
	messages <- ""

	// wait for everything to finish
	<-quit
	println("done")
}
