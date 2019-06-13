// Examples Test Cases
// 1           | 0 | 1
// 1           | 1 | 1
// 1           | 2 | 1
// 1->2        | 0 | 1->2
// 1->2        | 1 | 1->2
// 1->2        | 2 | 2->1
// 1->2->3     | 2 | 2->1->3
// 1->2->3->4  | 3 | 3->2->1->4
// 1->2->3->4  | 4 | 4->3->2->1
// 1->2->3->4  | 5 | 3->2->1->4
// 1->2->3->4  | 6 | 2->1->3->4

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
    if k <= 1 || head == nil {
        return head
    }
    
    dummy := &ListNode{-1, head}
    stack := make([]*ListNode, k)
    runner := 0
    start, finish := head, dummy
    for nil != start {
        stack[runner] = start
        runner++
        start = start.Next
        if runner == k {
            for runner > 0 {
                finish.Next = stack[runner-1]
                finish = finish.Next
                runner--
            }
        }
    }
    for i := 0; i < runner; i++ {
        finish.Next = stack[i]
        finish = finish.Next
    }
    finish.Next = nil
    
    return dummy.Next
}