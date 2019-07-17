/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    prev := math.MinInt64
    stack := []*TreeNode{}
    cur := root
    for len(stack) > 0 || cur != nil {
        if cur == nil {
            cur = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if prev >= cur.Val {
                return false
            }
            prev = cur.Val
            cur = cur.Right
        } else {
            stack = append(stack, cur)
            cur = cur.Left
        }
    }
    
    return true
}