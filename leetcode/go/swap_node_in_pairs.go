/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    temp := *head.Next
    temp.Next = head
    head.Next = swapPairs(head.Next.Next)
    return &temp
}