package main

import "fmt"

// Question

// Given an array of non-negative integers
// you are initially positioned at the first index of the array.
// Each element in the array represents your MAXimum jump length at that position.
// Determine if you are able to reach the last index.

// https://en.wikipedia.org/wiki/Dynamic_programming

// Input: [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

// Input: [3,2,1,0,4]
// Output: false
// Explanation: You will always arrive at index 3 no matter what. Its maximum
// jump length is 0, which makes it impossible to reach the last index.

// Solution
// This is a dynamic programming1 question.
// Usually, solving and fully understanding a dynamic programming problem is a 4 step process:
// 1 - Start with the recursive backtracking solution
// 2 - Optimize by using a memoization table (top-down dynamic programming)
// 3 - Remove the need for recursion (bottom-up dynamic programming)
// 4 - Apply final tricks to reduce the time / memory complexity

func canJumpG(nums []int) bool {
	lastPos := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Println(i, i+nums[i], lastPos)
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}

func main() {
	// fmt.Println(canJump([]int{}))
	fmt.Println(canJumpG([]int{1, 1, 3, 0, 5}))
	fmt.Println(canJumpG([]int{1, 0, 3, 0, 5}))
	fmt.Println(canJump([]int{1, 1, 3, 0, 5}))
	fmt.Println(canJump([]int{1, 0, 3, 0, 5}))
}

func canJump(nums []int) bool {
	for i := len(nums) - 2; i >= 0; i-- {
		// if we encounter 0, we're screwed
		fmt.Printf("i=%v || nums[i]=%v || i-1=%v\n", i, nums[i], i-1)

		if nums[i] != 0 {
			fmt.Println("continue..")
			continue
		}
		fmt.Println("if nums[i] lands on a 0 make sure we can leap it...")
		fmt.Println("----------")
		j := i - 1
		// for ; means like forever
		for ; j >= 0; j-- {
			fmt.Printf("j=%v || nums[j]=%v || i-j=%v\n", j, i-j, nums[j])
			if i-j < nums[j] {
				fmt.Println("----------")
				i = j
				break
			}
		}

		if j == -1 {
			fmt.Println("nope")
			return false
		}
	}
	return true
}
