/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// use bfs and queue

func largestValues(root *TreeNode) []int {
    var res []int
    if root == nil {
        return res
    }
    
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        qLen := len(queue)
        var largestVal int
        for i := 0; i < qLen; i++ {
            node := queue[0]
            queue = queue[1:]
            if node.Val > largestVal || i == 0 {
                largestVal = node.Val
            }
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        res = append(res, largestVal)
    }
    return res
}