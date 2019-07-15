/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
    var res int
    if root == nil {
        return res
    }
    
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        root = queue[0]
        queue = queue[1:]
        if root.Right != nil {
            queue = append(queue, root.Right)
        }
        if root.Left != nil {
            queue = append(queue, root.Left)
        }
    }
    return root.Val
}