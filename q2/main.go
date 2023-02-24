package main

import "fmt"

func main() {
	// Create a new buffered channel of strings with capacity 1.
	messages := make(chan string, 1)

	// Start a new goroutine that sends "contact" on the channel.
	go func() {
		messages <- "contact"
	}()

	// Wait for a message to be received on the channel and print it.
	msg := <-messages
	fmt.Println(msg)
}
