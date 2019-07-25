// recursive, DFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    leaves1, leaves2 := []int{}, []int{}    
    dfs(root1, &leaves1)
    dfs(root2, &leaves2)
    return  reflect.DeepEqual(leaves1, leaves2)
}

func dfs(node *TreeNode, leafValues*[]int) {
    if node != nil {
        if node.Left == nil && node.Right == nil {
            *leafValues = append(*leafValues, node.Val)
            return
        }
        dfs(node.Left, leafValues)
        dfs(node.Right, leafValues)
    }
}