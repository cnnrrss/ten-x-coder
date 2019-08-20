package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Heat *ListNode
	Tail *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	dummyNode := l3
	carry := 0
	for {
		sum := l1.Val + l2.Val + carry
		l3.Val = sum % 10
		carry = sum / 10
		if l1.Next == nil && l2.Next == nil {
			if carry == 1 {
				l3.Next = &ListNode{1, nil}
			}
			return dummyNode
		}
		l3.Next = &ListNode{}
		l3 = l3.Next
		if l1.Next == nil {
			l1.Val = 0
		} else {
			l1 = l1.Next
		}
		if l2.Next == nil {
			l2.Val = 0
		} else {
			l2 = l2.Next
		}

	}
}

func setupLists() (l1, l2 *ListNode) {
	l1 = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	return l1, l2
}

func main() {
	l1, _ := setupLists()
	l3 := addTwoNumbers(l1, &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			// Next: &ListNode{
			// 	Val: 7,
			// },
		},
	})
	fmt.Println(l3.Val, l3.Next.Val, l3.Next.Next.Val)
}
