package datastructures

import "fmt"

// Tree is a Binary Tree
type Tree struct {
	Root *TreeNode
}

// TreeNode is a Node of a Binary Tree
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// NewTreeNode is the constructor for a TreeNode
func NewTreeNode(v int, l, r *TreeNode) *TreeNode {
	return &TreeNode{Val: v, Left: l, Right: r}
}

// Insert add a new val to Tree
func (t *Tree) Insert(val int) {
	n := NewTreeNode(val, nil, nil)
	if t.Root == nil {
		t.Root = n
	} else {
		insertNode(t.Root, n)
	}
}

func insertNode(parent, child *TreeNode) {
	if child.Val < parent.Val {
		if parent.Left == nil {
			parent.Left = child
		} else {
			insertNode(parent.Left, child)
		}
	} else {
		if parent.Right == nil {
			parent.Right = child
		} else {
			insertNode(parent.Right, child)
		}
	}
}

// PrintValue prints the value of the TreeNode
func (n *TreeNode) PrintValue() {
	fmt.Printf("%d ", n.Val)
}

// String prints a visual representation of the tree
func (t *Tree) String() {
	fmt.Println("------------------------------------------------")
	stringify(t.Root, 0)
	fmt.Println("------------------------------------------------")
}

// internal recursive function to print a tree
func stringify(n *TreeNode, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.Left, level)
		fmt.Printf(format+"%d\n", n.Val)
		stringify(n.Right, level)
	}
}

// BuildTreeFromInts creates a Binary Tree from a slice of integers
func BuildTreeFromInts(nodeList []int) *Tree {
	tree := &Tree{}
	for _, node := range nodeList {
		tree.Insert(node)
	}

	return tree
}
