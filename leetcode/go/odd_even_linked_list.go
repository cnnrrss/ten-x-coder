/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    // Two pointers trick
    odd, even := head, head.Next
    evenHead := even
    
    for {
        if even == nil || even.Next == nil {
            break
        } else {
            odd.Next = even.Next
            odd = odd.Next
            even.Next = odd.Next
            even = even.Next
        }
    }
    
    odd.Next = evenHead
    return head
}