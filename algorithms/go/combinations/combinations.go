package main

import (
	"fmt"
)

func main() {
	// 0 - 4, print 3
	comb(5, 3)
}

func comb(n, m int) {
	s := make([]int, m)
	last := m - 1

	emit := func(c []int) {
		fmt.Println(c)
	}

	var rc func(int, int)
	rc = func(i, next int) {
		for j := next; j < n; j++ {
			s[i] = j
			if i == last {
				emit(s)
			} else {
				rc(i+1, j+1)
			}
		}
		return
	}

	rc(0, 0)
}
