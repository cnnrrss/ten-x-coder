/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    ans := []int{}
    if root == nil {
        return ans
    }
    stack := &NodeStack{}
    stack.push(root)
    for {
        curr := stack.pop()
        if curr == nil {
            break
        }
        ans = append(ans, curr.Val)
        if curr.Right != nil {
            stack.push(curr.Right)
        }
        if curr.Left != nil {
            stack.push(curr.Left)
        }
    }
    return ans
}

type NodeStack struct {
    Items []*TreeNode
}

func (s *NodeStack) pop() *TreeNode {
    length := len(s.Items)
    if length == 0 {
        return nil
    }
    
    item, items := s.Items[length-1], s.Items[0:length-1]
    s.Items = items
    return item
}

func (s *NodeStack) push(n *TreeNode) {
    s.Items = append(s.Items, n)
}