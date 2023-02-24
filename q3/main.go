package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	complaints := make(chan string)
	go patient(complaints)
	go doctor(complaints)
	time.Sleep(5 * time.Second)
	fmt.Println("Time's up, the clinic is closing")
}

func patient(complaints chan string) {
	for {
		complaints <- fmt.Sprintf("Patient: my leg hurts (%v)", time.Now().Format(time.Stamp))
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}
}

func doctor(complaints chan string) {
	for {
		complaint := <-complaints
		fmt.Printf("Doctor: received complaint %q (%v)\n", complaint, time.Now().Format(time.Stamp))
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		diagnosis := fmt.Sprintf("Doctor: diagnosis for %q is %q (%v)", complaint, "rest and painkillers", time.Now().Format(time.Stamp))
		fmt.Println(diagnosis)
	}
}
