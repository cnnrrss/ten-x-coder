package main

import ds "github.com/cnnrrss/ten-x-coder/algorithms/go/datastructures"

func inOrderTraversalRecursion(root *ds.TreeNode) {
	if root == nil {
		return
	}
	inOrderTraversalRecursion(root.Left)
	root.PrintValue()
	inOrderTraversalRecursion(root.Right)
}
