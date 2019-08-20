package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// The primary mechanism for managing state in Go is communication over channels.
// We saw this for example with worker pools. There are a few other options for
// managing state though. Here we’ll look at using the sync/atomic package for
// atomic counters accessed by multiple goroutines.

// In order to safely use the counter while it’s still being updated by other goroutines,
// we extract a copy of the current value into opsFinal via LoadUint64. As above we need
// to give this function the memory address &ops from which to fetch the value.

func main() {
	var ops uint64
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond * 1)
			}
		}()
	}
	time.Sleep(time.Millisecond * 100)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
