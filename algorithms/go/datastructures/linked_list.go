package datastructures

import "fmt"

// Ele is a node element in a linked list
type Ele struct {
	Data interface{}
	Next *Ele
}

// Append adds a new ele to a linked list
func (e *Ele) Append(data interface{}) *Ele {
	if e.Next == nil {
		e.Next = &Ele{data, nil}
	} else {
		tmp := &Ele{data, e.Next}
		e.Next = tmp
	}
	return e.Next
}

// String returns the current element
// of a linked list as a string
func (e *Ele) String() string {
	return fmt.Sprintf("Ele: %v", e.Data)
}
