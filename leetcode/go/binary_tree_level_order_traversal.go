// Implemented BFS using Queue

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Queue interface {
    Enqueue(node interface{})
    Dequeue() interface{}
    Size() int
    IsEmpty() bool
}

type NodeQueue struct {
    nodes []*TreeNode
}

func (s *NodeQueue) Enqueue(node *TreeNode) {
    s.nodes = append(s.nodes, node)
}

func (s *NodeQueue) Dequeue() *TreeNode {
    if s.IsEmpty() {
        return nil
    }
    
    node, nodes := s.nodes[0], s.nodes[1:]
    s.nodes = nodes
    return node
}

func (s *NodeQueue) Size() int {
    return len(s.nodes)
}

func (s *NodeQueue) IsEmpty() bool {
    return len(s.nodes) == 0
}


func levelOrder(root *TreeNode) [][]int {
    var res [][]int
    if root == nil {
        return res
    }
    
    queue := &NodeQueue{}
    queue.Enqueue(root)
    for !queue.IsEmpty() {
        qLen := queue.Size()
        var levelRes []int
        for i := 0; i < qLen; i++ {
            node := queue.Dequeue()
            levelRes = append(levelRes, node.Val)
            
            if node.Left != nil {
                queue.Enqueue(node.Left)   
            }
            if node.Right != nil {
                queue.Enqueue(node.Right)   
            }
        }
        
        if len(levelRes) > 0 {
            res = append(res, levelRes)   
        }
    }
    return res
}