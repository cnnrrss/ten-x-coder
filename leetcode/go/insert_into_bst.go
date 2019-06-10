/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// First recursive approach
 func insertIntoBSTRecursice(root *TreeNode, val int) *TreeNode {
     if root == nil {
         return &TreeNode{Val: val}
     }
     if root.Val > val {
         root.Left = insertIntoBST(root.Left, val)
     } else {
         root.Right = insertIntoBST(root.Right, val)
     }
     return root
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
    node := root
    for {
        if node == nil {
            return &TreeNode{Val: val}
        }
        
        if node.Val > val {
            if node.Left == nil {
                node.Left = &TreeNode{Val: val}
                return root
            }
            node = node.Left
        } else {
            if node.Right == nil {
                node.Right = &TreeNode{Val: val}
                return root
            }
            node = node.Right
        }
    }
}