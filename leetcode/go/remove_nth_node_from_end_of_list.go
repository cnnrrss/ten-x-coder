/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil{
        return nil
    }
    
    if n == 0 && head.Next != nil {
        return head.Next
    }
    
    // ans to be returned
    ans := &ListNode{Val: 0}
    ans.Next = head
    // two pointers
    first, second := ans, ans
    
    // advance second pointer so it is n nodes apart from first
    for i := 1; i <= n + 1; i++ {
        second = second.Next
    }
    
    // Move running ptr (second) to end
    for second != nil {
        first = first.Next
        second = second.Next
    }

    first.Next = first.Next.Next
        
    return ans.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
    if head == nil {
        return nil
    }

    // two pointers
    slow, fast := head, head
    
    // advance second pointer so it is n nodes apart from first
    for i := 0; i < n; i++ {
        if fast.Next == nil {
            return head.Next
        }
        fast = fast.Next
    }
    
    // Move running ptr (second) to end
    for fast.Next != nil {
        fast, slow = fast.Next, slow.Next
    }

    slow.Next = slow.Next.Next
        
    return head
}