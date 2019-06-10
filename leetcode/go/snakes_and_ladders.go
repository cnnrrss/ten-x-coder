package main

import "fmt"

func main() {
	board := [][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1},
	}

	fmt.Println(snakesAndLadders(board))
}

func snakesAndLadders(board [][]int) int {
	n := len(board)
	dist := map[int]int{}
	dist[1] = 0

	queue := []int{1}
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		if s == n*n {
			return dist[s]
		}

		for s2 := s + 1; s2 <= min(s+6, n*n); s2++ {
			if s2 >= min(s+6, n*n) {
				break
			}
			r, c := getPosition(s2+1, n)
			if board[r][c] != -1 {
				s2 = board[r][c]
			}
			if _, ok := dist[s2+1]; !ok {
				dist[s2] = dist[s] + 1
				queue = append(queue, s2)
			}
		}
	}

	return -1
}

func getPosition(s, n int) (x int, y int) {
	s--
	x, y = s/n, s%n
	if x%2 == 1 {
		y = n - y - 1
	}
	x = n - x - 1
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
