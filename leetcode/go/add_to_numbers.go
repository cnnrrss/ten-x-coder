/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    l3 := &ListNode{}
    res := l3
    carry, sum := 0, 0
    for {
        sum = l1.Val + l2.Val + carry
        l3.Val = sum % 10
        carry = sum / 10

        if l1.Next == nil && l2.Next == nil {
            if carry == 1 {
                l3.Next = &ListNode{Val: 1, Next: nil}
            }
            return res
        }
        
        if l1.Next != nil {
            l1 = l1.Next
        } else {
            l1.Val = 0
        }
        
        if l2.Next != nil {
            l2 = l2.Next
        } else {
            l2.Val = 0
        }
        
        l3.Next = &ListNode{}
        l3 = l3.Next
    }

}