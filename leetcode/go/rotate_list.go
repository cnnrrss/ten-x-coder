/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 * 1) Validate the input
 * 2) Get the length of the list (not stored)
 * 3) 
 */

func rotateRight(head *ListNode, k int) *ListNode {
    if k == 0 || head == nil || head.Next == nil {
        return head
    }
    
    var runner, newList, temp *ListNode
    runner = head
    length := 1 // start at 1 because we terminate we know head >= 1
    
    for runner.Next != nil {
        runner = runner.Next
        length++
    }
    // return early if length = k
    if k == length {
        return head
    }
    // if k > length find mod so we don't do unnecessary work (BUD)
    if k > length {
        k = k % length
        if k == 0 {
            return head
        }
    }
    
    // reset runner, run again
    runner = head
    
    // stop k-1 nodes from end
    for i := 0; i < length-k-1; i++ {
        runner = runner.Next
    }

    // set head of new List kth node from end of orig
    newList = runner.Next
    // set runner kth node
    runner.Next = nil
    temp = newList
    
    for temp.Next != nil {
        temp = temp.Next
    }
    temp.Next = head
    
    return newList
}