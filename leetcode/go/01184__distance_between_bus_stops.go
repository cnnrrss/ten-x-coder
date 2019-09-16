package main

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	trip, total := 0, 0
	if start > destination {
		start, destination = destination, start
	}

	for i, leg := range distance {
		total += leg
		if start <= i && i < destination {
			trip += leg
		}
	}
	return min(trip, total-trip)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
