package main

/**
198. House Robber
You are a professional robber planning to rob houses along a street.
Each house has a certain amount of money stashed, the only constraint
stopping you from robbing each of them is that adjacent houses have
security system connected and it will automatically contact the police
if two adjacent houses were broken into on the same night.
Given a list of non-negative integers representing the amount of money
of each house, determine the maximum amount of money you can rob tonight
without alerting the police.

Credits:Special thanks to @ifanchu for adding this problem and creating
all test cases. Also thanks to @ts for adding additional test cases.

[2, 4, 1, 6]
*/

import "fmt"

func rob(nums []int) int {
	var a, b int
	for i, v := range nums {
		if i%2 == 0 {
			a = max(a+v, b)
		} else {
			b = max(b+v, a)
		}
	}
	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	sum := rob([]int{1, 2, 3, 5, 4, 9, 3, 2, 10})
	fmt.Println(sum)
}
