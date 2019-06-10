/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var sum int
func convertBST(root *TreeNode) *TreeNode {    
    sum = 0
    traverse(root)
    return root
}

func traverse(root *TreeNode)  {
    if root == nil {
        return
    }
    
    // In order beyotcch
    traverse(root.Right)

    sum += root.Val
    root.Val = sum
    
    traverse(root.Left)
}