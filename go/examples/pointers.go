package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int
	var t bool
	x = 5
	t = false
	y := &x
	f := &t
	fmt.Println(reflect.TypeOf(y), y, &y, *y) // *int, locX: 0x414020, locY: 0x40c130, 5
	fmt.Println(reflect.TypeOf(f), f, &f, *f) // *bool, locT: 0x414024, locF: 0x40c138, false
	// you cannot do *x, *t because variables x, t are not of type pointer
	// ./prog.go:17:14: invalid indirect of x (type int)
	// ./prog.go:17:18: invalid indirect of t (type bool)
	fmt.Println(&x, &t) // this references the pointer and displays the memory address
}
