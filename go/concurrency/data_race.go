package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

}

func dataRace() {
	var data int
	go func() { data++ }() // 1
	if data == 0 {         // 2
		fmt.Printf("what is data anyway? %d\n", data) // 3
	}
}

// have we fixed it by sleeping? Nah bro!
func dataRace2() {
	var data int
	go func() { data++ }() // 1
	time.Sleep(1 * time.Second)
	if data == 0 { // 2
		fmt.Printf("what is data anyway? %d\n", data) // 3
	}
}

// Solved our data race, but we haven’t actually
// solved our race condition! Order of ops is still
// non deterministic which occurs first? 1, 2, or 3
func dataRaceLocked() {
	var data int
	var memAccess sync.Mutex
	go func() {
		memAccess.Lock()
		data++ // 1
		memAccess.Unlock()
	}()
	memAccess.Lock()
	if data == 0 { // 2
		fmt.Printf("what is data anyway? %d\n", data)
	} else {
		fmt.Printf("data is chill %d\n", data) // 3
	}
	memAccess.Unlock()
}
