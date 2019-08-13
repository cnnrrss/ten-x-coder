package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	// if only 1 interval, return it
	if len(intervals) <= 1 {
		return intervals
	}

	// sort.Slice, sorts the slice of intervals given the provided less func
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// make new slice of intervals for result
	res := make([][]int, 0, len(intervals))

	// set initial last interval as first sorted interval
	last := intervals[0]
	// loop thru intervals
	for i := 1; i < len(intervals); i++ {
		// if next interval begins before the last end
		if intervals[i][0] <= last[1] {
			// new end time is max of both end times
			last[1] = max(last[1], intervals[i][1])
		} else {
			// if no overlap, append last interval to result
			res = append(res, last)
			// update last interval to current one
			last = intervals[i]
		}
	}
	res = append(res, last)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	intervals := [][]int{{1, 3}, {15, 18}, {2, 6}, {14, 17}, {8, 10}}
	fmt.Println(intervals)
	intervals = merge(intervals)
	fmt.Println(intervals)
}
