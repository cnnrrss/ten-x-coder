/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    for root != nil {
        if root.Val > p.Val && root.Val > q.Val {
            root = root.Left
        } else if root.Val < p.Val  && root.Val < q.Val {
            root = root.Right
        } else {
            return root
        }
    }  
    return nil
}