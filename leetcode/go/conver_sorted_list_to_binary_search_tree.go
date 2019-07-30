/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
    return helper(head, nil)
}

func helper(head, tail *ListNode) *TreeNode {
    if head == tail {
        return nil
    }
    
    if head.Next == tail {
        return &TreeNode{Val: head.Val}
    }

    slow, fast := head, head
    for fast != tail && fast.Next != tail {
        fast = fast.Next.Next
        slow = slow.Next
    }
    
    mid := slow

    return &TreeNode{
        Val: mid.Val,
        Left: helper(head, mid),
        Right: helper(mid.Next, tail),
    }
}