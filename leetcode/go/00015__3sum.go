import "sort"

// make zero, zero is target
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	solution := [][]int{}
	// sorting can help make this O(n^2)
	sort.Ints(nums)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] { // if duplicate numbers, continue
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			tempSum := nums[i] + nums[l] + nums[r]
			switch {
			case tempSum > 0:
				r--
			case tempSum < 0:
				l++
			default:
				solution = append(solution, []int{nums[i], nums[l], nums[r]})
				l, r = next(nums, l, r)
			}
		}
	}
	return solution
}

func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}
	return l, r
}

// Algorithm:
// - Sort
// - Plant a at index 0, left = index+1, right = len(arr)-1
// - move left and right into each other until pair is found
// - skip duplicates
// - move left or right depending on whether sum if > or < 0
// - advance after solution with nextFunc
//      - if the next left, or next right is the same, skip
//      - else move left, and right (move both because we can't have a dupe solution)

// input
// [-1, 0, 1, 2, -1, -4],

// sort
//   i  l             r
// [-4, -1, -1, 0, 1, 2]

// skip index 2 because dupe
//   i          l     r
// [-4, -1, -1, 0, 1, 2]

// move l again because 3sum still < 0
//   i             l  r
// [-4, -1, -1, 0, 1, 2]

// iter
//       i   l        r
// [-4, -1, -1, 0, 1, 2] = valid sol

// // next move both
//           i  l  r
// [-4, -1, -1, 0, 1, 2]

// move r because 3sum > 0
//           i  l  r
// [-4, -1, -1, 0, 1, 2] = valid sol
// end
