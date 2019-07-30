/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func partition(head *ListNode, x int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    beforeList, afterList := &ListNode{}, &ListNode{}
    before, after := beforeList, afterList

    for head != nil {
        if head.Val < x {
            before.Next = head
            before = before.Next
        } else {
            after.Next = head
            after = after.Next
        }
        head = head.Next
    }

    before.Next = afterList.Next
    after.Next = nil
    return beforeList.Next
}
