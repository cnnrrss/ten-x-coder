package main

import (
	"fmt"

	ds "github.com/cnnrrss/ten-x-coder/algorithms/go/datastructures"
)

func preOrderTraversalRecursion(root *ds.TreeNode) {
	if root == nil {
		return
	}
	root.PrintValue()
	preOrderTraversalRecursion(root.Left)
	preOrderTraversalRecursion(root.Right)
}

func preOrderTraversalIterative(root *ds.TreeNode) {
	if root == nil {
		return
	}

	stack := &ds.NodeStack{}
	stack.Push(root)
	for {
		temp := stack.Pop()
		if temp == nil {
			break
		}
		temp.PrintValue()
		// trick here is to push right first in iterative preOrder
		// because a stack is LIFO datastructure
		if temp.Right != nil {
			stack.Push(temp.Right)
		}
		if temp.Left != nil {
			stack.Push(temp.Left)
		}
	}
}

// main driver function for traversal
func main() {
	nodeList := []int{100, 50, 150, 25, 75, 125, 175, 110}
	t := ds.BuildTreeFromInts(nodeList)

	preOrderTraversalRecursion(t.Root)
	fmt.Printf("\n")

	preOrderTraversalIterative(t.Root)
	fmt.Printf("\n")

	preOrderTraversalIterative(t.Root)
}
