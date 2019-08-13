package main

import "sort"

func canAttendMeetings(intervals [][]int) bool {
	// if only 1 meeting then you can attend all
	if len(intervals) <= 1 {
		return true
	}

	// sort the meetings
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// for all meetings
	for i := 0; i < len(intervals)-1; i++ {
		// if next meeting starts before this one ends
		if intervals[i+1][0] < intervals[i][1] {
			return false
		}
	}
	return true
}
