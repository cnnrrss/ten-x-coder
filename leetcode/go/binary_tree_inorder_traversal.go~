/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
    resp := &[]int{}
    recursive(root, resp)
    return *resp
}

func recursive(root *TreeNode, resp *[]int) {
    if root != nil {
        if root.Left != nil {
            recursive(root.Left, resp)
        }
        *resp = append(*resp, root.Val)
        if root.Right != nil {
            recursive(root.Right, resp)
        }
    }
}