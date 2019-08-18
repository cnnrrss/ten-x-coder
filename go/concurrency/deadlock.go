package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
}

type value struct {
	mu    sync.Mutex
	value int
}

func deadlockMeBaby() {
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()                // 1
		defer v1.mu.Unlock()        // 2
		time.Sleep(2 * time.Second) // 3
		v2.mu.Lock()
		defer v2.mu.Unlock()
		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}
	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}

// 1 Here we attempt to enter the critical section for the incoming value.
// 2 Here we use the defer statement to exit the critical section before printSum returns.
// 3 Here we sleep for a period of time to simulate work (and trigger a deadlock).
//  this code emits: {fatal error: all goroutines are asleep - deadlock!}
// see figure 1.1
// We have created two gears that cannot turn together:
// Our first call to printSum locks a and then attempts to lock b
// In the meantime our second call to printSum, locked b and attempts to lock a.
// Both goroutines wait infinitely on each other.
