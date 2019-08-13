package datastructures

// NodeStack is a adt stack of TreeNodes
type NodeStack struct {
	items []*TreeNode
}

// Push ...
func (s *NodeStack) Push(n *TreeNode) {
	s.items = append(s.items, n)
}

// Pop ...
func (s *NodeStack) Pop() *TreeNode {
	length := len(s.items)
	if length == 0 {
		return nil
	}

	// remember this
	item, items := s.items[length-1], s.items[0:length-1]
	s.items = items
	return item
}

// Top returns the top TreeNode from the NodeStack
func (s *NodeStack) Top() *TreeNode {
	return s.items[len(s.items)-1]
}
