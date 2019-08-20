package main

import "fmt"

// When using channels as function parameters,
// you can specify if a channel is meant to
// only send or receive values. This specificity
// increases the type-safety of the program.

// This ping function only accepts a channel for sending values.
// It would be a compile-time error to try to receive on this channel.
func ping(healthcheckz chan<- string, msg string) {
	healthcheckz <- msg
}

// The relay function accepts one channel for receives (healthcheckz)
//  and a second for sends (relays).
func relay(healthcheckz <-chan string, relays chan<- string) {
	// msg := <-healthcheckz
	// relays <- msg

	// one liner
	relays <- <-healthcheckz
}

func main() {
	healthcheckz := make(chan string, 1)
	relays := make(chan string, 1)
	ping(healthcheckz, "passed message")
	relay(healthcheckz, relays)
	fmt.Println(<-relays)
}
