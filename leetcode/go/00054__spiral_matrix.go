
// Time Complexity: O(N), where N is the total number of elements in the input matrix.
// We add every element in the matrix to our final answer.

// Space Complexity: O(N), the information stored in ans.

// [[1, 1, 1, 1, 1, 1, 1], top: c from c1 ... c2
//  [1, 2, 2, 2, 2, 2, 1], right: r from r1+1 ... r2
//  [1, 2, 3, 3, 3, 2, 1], bottom: c from c2+1 ... c1+1
//  [1, 2, 2, 2, 2, 2, 1], left r from r2+1 ... r1+1
//  [1, 1, 1, 1, 1, 1, 1]]

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	r1, r2, c1, c2 := 0, len(matrix)-1, 0, len(matrix[0])-1
	ans := []int{}
	for r1 <= r2 && c1 <= c2 {
		for c := c1; c <= c2; c++ {
			ans = append(ans, matrix[r1][c])
		}
		for r := r1 + 1; r <= r2; r++ {
			ans = append(ans, matrix[r][c2])
		}
		if r1 < r2 && c1 < c2 {
			for c := c2 - 1; c > c1; c-- {
				ans = append(ans, matrix[r2][c])
			}
			for r := r2; r > r1; r-- {
				ans = append(ans, matrix[r][c1])
			}
		}
		r1++
		r2--
		c1++
		c2--
	}
	return ans
}