/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    ans := 0
    
    var rootToLeaf func(root *TreeNode, temp int) 
    rootToLeaf = func(root *TreeNode, temp int) {
        if root == nil {
            return
        }

        // This is the B&B
        temp = temp*10 + root.Val

        if root.Left == nil && root.Right == nil {
            ans += temp
            return
        }
        
        // recursive call on the leaves
        rootToLeaf(root.Left, temp)
        rootToLeaf(root.Right, temp)
    }
    
    rootToLeaf(root, 0)
    
    return ans
}