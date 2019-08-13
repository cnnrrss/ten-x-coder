/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Step through LL and insert into arr


func nextLargerNodes(head *ListNode) []int {
    if head == nil {
        return nil
    }
    
    arr, size := listNodeToArray(head)
    // [2,1,5] | 3 | - 1
    stack, top := make([]int, 0, size), -1
    ans := make([]int, size)
    for i:=size-1; i >= 0; i-- {
        for top >= 0 && stack[top] <= arr[i] {
            top--
        }
        
        stack = stack[:top+1]
        if top >= 0 {
           ans[i] = stack[top]
        }
        
        top++
        stack = append(stack, arr[i])
    }
    return ans
}

func listNodeToArray(head *ListNode) ([]int, int) {
    var arr = make([]int, 0)
    
    for head != nil {
        arr = append(arr, head.Val)
        head = head.Next
    }
    
    return arr, len(arr)
}