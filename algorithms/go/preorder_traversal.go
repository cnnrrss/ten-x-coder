package main

import "fmt"

type node struct {
	left  *node
	right *node
	val   int
}

type tree struct {
	root *node
}

func (t *tree) insert(val int) {
	n := newNode(val, nil, nil)
	if t.root == nil {
		t.root = n
	} else {
		insertNode(t.root, n)
	}
}

func insertNode(parent, child *node) {
	if child.val < parent.val {
		if parent.left == nil {
			parent.left = child
		} else {
			insertNode(parent.left, child)
		}
	} else {
		if parent.right == nil {
			parent.right = child
		} else {
			insertNode(parent.right, child)
		}
	}
}

func newNode(v int, l, r *node) *node {
	return &node{val: v, left: l, right: r}
}

func (n *node) printValue() {
	fmt.Printf("%d ", n.val)
}

type nodeStack struct {
	items []*node
}

func (s *nodeStack) push(n *node) {
	s.items = append(s.items, n)
}

func (s *nodeStack) pop() *node {
	length := len(s.items)
	if length == 0 {
		return nil
	}

	// remember this
	item, items := s.items[length-1], s.items[0:length-1]
	s.items = items
	return item
}

func (s *nodeStack) Top() *node {
	return s.items[len(s.items)-1]
}

func buildTree(nodeList []int) *tree {
	tree := &tree{}
	for _, node := range nodeList {
		tree.insert(node)
	}

	return tree
}

func preOrderTraversalRecursion(root *node) {
	if root == nil {
		return
	}
	root.printValue()
	preOrderTraversalRecursion(root.left)
	preOrderTraversalRecursion(root.right)
}

func preOrderTraversal(root *node) {
	if root == nil {
		return
	}

	stack := &nodeStack{}
	stack.push(root)
	for {
		temp := stack.pop()
		if temp == nil {
			break
		}
		temp.printValue()
		// trick here is to push right first in iterative preOrder
		// because a stack is LIFO datastructure
		if temp.right != nil {
			stack.push(temp.right)
		}
		if temp.left != nil {
			stack.push(temp.left)
		}
	}
}

func inOrderTraversalRecursion(root *node) {
	if root == nil {
		return
	}
	inOrderTraversalRecursion(root.left)
	root.printValue()
	inOrderTraversalRecursion(root.right)
}

func postOrderTraversalRecursion(root *node) {
	if root == nil {
		return
	}
	postOrderTraversalRecursion(root.left)
	postOrderTraversalRecursion(root.right)
	root.printValue()
}

func BFS() {
	return
}

func DFS() {
	return
}

// String prints a visual representation of the tree
func (t *tree) String() {
	fmt.Println("------------------------------------------------")
	stringify(t.root, 0)
	fmt.Println("------------------------------------------------")
}

//
func main() {
	nodeList := []int{100, 50, 150, 25, 75, 125, 175, 110}
	// nodeList := []int{10, 20, 30, 40, 50, 60, 70}
	t := buildTree(nodeList)
	preOrderTraversalRecursion(t.root)
	fmt.Printf("\n")
	preOrderTraversal(t.root)
	fmt.Printf("\n")
	preOrderTraversal(t.root)
	// t.postOrderTraversal()
	// t.inOrderTraversal()
	// t.BFS()
	// t.DFS()
	// t.String()
}

// internal recursive function to print a tree
func stringify(n *node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.val)
		stringify(n.right, level)
	}
}
