package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	BFS(root)
	invertBT(root)
	BFS(root)
}

func invertBT(root *TreeNode) {
	if root == nil {
		return
	}

	root.Left, root.Right = root.Right, root.Left

	if root.Left != nil {
		invertBT(root.Left)
	}

	if root.Right != nil {
		invertBT(root.Right)
	}
	return
}

func BFS(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%d,", node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	fmt.Println()
}

func dfsRecursive(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	dfsRecursive(root.Left) 
	dfsRecursive(root.Right
	return
}
//		1
//     /  \
// 	  3     10
//   / \   / \
//  7   5 6   9
//
// [1, 3, 10, 7, 5, 6, 9]
//
//		1
//     /  \
// 	  10    3
//   / \   / \
//  9   6 5   7
//
// [1, 10, 3, 9, 6, 5, 7]
//
