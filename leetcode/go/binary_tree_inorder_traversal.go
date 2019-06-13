/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversalIterative(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    ans := []int{}
    stack := &NodeStack{}
    curr := root
    
    for curr != nil || len(stack.Items) > 0 {    
        for curr != nil {
            stack.push(curr)
            curr = curr.Left
        }
        
        curr = stack.pop()
        ans = append(ans, curr.Val)
        curr = curr.Right

    }
    return ans
}

type NodeStack struct {
    Items []*TreeNode
}

func (s *NodeStack) pop() *TreeNode {
    if len(s.Items) == 0 {
        return nil
    }
    
    node, items := s.Items[len(s.Items)-1], s.Items[0:len(s.Items)-1]
    s.Items = items
    return node
}

func (s *NodeStack) push(node *TreeNode) {
    s.Items = append(s.Items, node)
}

/**
* Recursive Solution
*/
func inorderTraversalRecursive(root *TreeNode) []int {
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