package main

import "fmt"

func main() {
	// var t solver  // declare variable of solver type
	t := &towers{} // type towers must satisfy solver interface
	t.play(3)      // # of disks
}

// a towers of hanoi solver just has one method, play
type solver interface {
	play(int)
}

// towers type satisfyies the solver interface
type towers struct{}

// play towers of hanoi
func (t *towers) play(n int) {
	t.moveDisks(n, 1, 2, 3)
}

// recursive
func (t *towers) moveDisks(n, from, to, via int) {
	if n > 0 { // handle base case
		t.moveDisks(n-1, from, via, to)
		t.moveDisk(n, from, to)
		t.moveDisks(n-1, via, to, from)
	}
}

func (t *towers) moveDisk(n, from, to int) {
	fmt.Printf("move disk %d from rod %d to rod %d\n", n, from, to)
}
