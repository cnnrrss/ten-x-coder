func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}

	if n == 1 {
		return [][]int{{1}}
	}

	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	next := nextFunc(n)

	for i := 1; i <= n*n; i++ {
		x, y := next()
		ans[x][y] = i
	}

	return ans
}

func nextFunc(n int) func() (int, int) {
	up, dn, lt, rt := 0, n-1, 0, n-1
	x, y := 0, -1
	dx, dy := 0, 1
	return func() (int, int) {
		x += dx
		y += dy
		switch {
		case y+dy > rt:
			up++
			dx, dy = 1, 0
		case x+dx > dn:
			rt--
			dx, dy = 0, -1
		case y+dy < lt:
			dn--
			dx, dy = -1, 0
		case x+dx < up:
			lt++
			dx, dy = 0, 1
		}
		return x, y
	}
}

// [
//     [1, 2],
//     [4, 3]
// ]

// [
//  [ 1, 2, 3 ],
//  [ 8, 9, 4 ],
//  [ 7, 6, 5 ]
// ]

// [
//     [01, 02, 03, 04],
//     [12, 13, 14, 05],
//     [11, 16, 15, 06],
//     [10, 09, 08, 07],
// ]