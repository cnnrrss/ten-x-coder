package datastructures

import "fmt"

// DLLNode is a node in a Doubly Linked List
type DLLNode struct {
	string
	next, prev *DLLNode
}

// DLLList is a Doubly Linked List
type DLLList struct {
	head, tail *DLLNode
}

func (list *DLLList) insertTail(node *DLLNode) {
	if list.tail == nil {
		list.head = node
	} else {
		list.tail.next = node
	}
	node.next = nil
	node.prev = list.tail
	list.tail = node
}

func (list *DLLList) insertAfter(existing, insert *DLLNode) {
	insert.prev = existing
	insert.next = existing.next
	existing.next.prev = insert
	existing.next = insert
	if existing == list.tail {
		list.tail = insert
	}
}

func (list *DLLList) String() string {
	if list.head == nil {
		return fmt.Sprint(list.head)
	}
	r := "[" + list.head.string
	for p := list.head.next; p != nil; p = p.next {
		r += " " + p.string
	}
	return r + "]"
}

// func main() {
// 	dll := &DLLList{}
// 	fmt.Println(dll)
// 	a := &DLLNode{string: "A"}
// 	dll.insertTail(a)
// 	dll.insertTail(&DLLNode{string: "B"})
// 	fmt.Println(dll)
// 	dll.insertAfter(a, &DLLNode{string: "C"})
// 	fmt.Println(dll)

// 	// traverse from end to beginning
// 	fmt.Print("From tail:")
// 	for p := dll.tail; p != nil; p = p.prev {
// 		fmt.Print(" ", p.string)
// 	}
// 	fmt.Println("")
// }
