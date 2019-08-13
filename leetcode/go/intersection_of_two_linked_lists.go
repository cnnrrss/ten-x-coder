/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    
    curA, curB := headA, headB
    
    for curA != nil && curB != nil {
        curA = curA.Next
        curB = curB.Next
    }

    for curA != nil {
        headA = headA.Next
        curA = curA.Next
    }

    for curB != nil {
        headB = headB.Next
        curB = curB.Next
    }

    for headA != headB {
        headA = headA.Next
        headB = headB.Next
    }
    return headA
}