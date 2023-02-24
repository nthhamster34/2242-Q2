package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// playSports simulates a sports event by sleeping for a random duration
// between 1 and 3 seconds, and then printing a message indicating that
// the event has finished. The WaitGroup pointer wg is used to signal
// when the event has finished.
func playSports(event string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Starting %s event\n", event)

	// Simulate the event by sleeping for a random duration between 1 and 3 seconds.
	r := time.Duration(1 + rand.Intn(3))
	time.Sleep(r * time.Second)

	fmt.Printf("Finished %s event\n", event)
}

func main() {
	var wg sync.WaitGroup

	// Define a slice of event names.
	events := []string{"soccer", "basketball", "tennis", "volleyball", "hockey"}

	// For each event, add 1 to the WaitGroup and launch a goroutine that
	// calls playSports with the event name and the WaitGroup pointer.
	for _, event := range events {
		wg.Add(1)

		go playSports(event, &wg)
	}

	// Wait for all the goroutines to finish.
	wg.Wait()

	fmt.Println("All events finished")
}
