func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	m, n := len(image), len(image[0])
	prevColor := image[sr][sc]
	if newColor == prevColor {
		return image
	}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if image[r][c] == prevColor {
			image[r][c] = newColor
			if r >= 1 {
				dfs(r-1, c)
			}
			if r+1 < m {
				dfs(r+1, c)
			}
			if c >= 1 {
				dfs(r, c-1)
			}
			if c+1 < n {
				dfs(r, c+1)
			}
		}
	}
	dfs(sr, sc)

	return image
}