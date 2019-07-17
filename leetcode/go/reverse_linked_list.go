/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    var rev *ListNode
    for head != nil {
        head.Next, rev, head = rev, head, head.Next     
    }
    return rev
}