package main

import (
	"fmt"
)

func nextPermutation(nums []int) {
	// fmt.Printf("%v", nums)
	length := len(nums)

	for i := length - 1; i > 0; i-- {
		fmt.Printf("%v ->> ", i)
		minPost := i
		for j := i; j < length; j++ {
			if nums[j] <= nums[minPos] && nums[j] > nums[i-1] {
				minPos = j
			}
			fmt.Printf("%v\n", j)
		}
	}
	// fmt.Printf(" -> %v\n ", nums)
}

func main() {
	nextPermutation([]int{1, 2, 3})
	nextPermutation([]int{4, 3, 2, 1})
	// nextPermutation([]int{1, 3, 2})

	// [1,3,2]
	// [1,2,3]
	// [1,5,1]
	// [5,1,1]
	// [1,2,4,3]
	// [1,2,3,4]
	// [3,1,2]
}
