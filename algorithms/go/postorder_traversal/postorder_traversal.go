package main

import (
	"fmt"

	ds "github.com/cnnrrss/ten-x-coder/algorithms/go/datastructures"
)

func postOrderTraversalRecursion(root *ds.TreeNode) {
	if root == nil {
		return
	}
	postOrderTraversalRecursion(root.Left)
	postOrderTraversalRecursion(root.Right)
	root.PrintValue()
}

// main driver function for traversal
func main() {
	nodeList := []int{100, 50, 150, 25, 75, 125, 175, 110}
	t := ds.BuildTreeFromInts(nodeList)

	postOrderTraversalRecursion(t.Root)
	fmt.Printf("\n")

	postOrderTraversalRecursion(t.Root)
	fmt.Printf("\n")

	postOrderTraversalRecursion(t.Root)
}
