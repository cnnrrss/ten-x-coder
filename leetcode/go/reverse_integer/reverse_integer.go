package main

// import "fmt"

func reverse(x int) int {
	y := 0

	if 0 <= x && x < 10 {
		return x
	}

	// for _, i := range x {
	t := x % 10
	if t == 0 {
		// handle overflow
	}

	x = x / 10
	y = y*10 + t
	// }
	return x
}

// Given a 32-bit signed integer, reverse digits of an integer.
//
// Note:
// Assume we are dealing with an environment which could only
// store integers within the 32-bit signed integer range: [−231,  231 − 1].
// For the purpose of this problem, assume that your function
// returns 0 when the reversed integer overflows
func reverse2(x int) int {
	var ans, temp, sign int
	max := 1<<31 - 1

	if x >= 0 {
		sign = 1
	} else {
		sign = -1
		x = -x
	}

	for x >= 10 {
		temp = x % 10 // get last integer
		x = x / 10    // remove last integer
		ans = ans*10 + temp
	}

	temp = x % 10

	if temp > 0 {
		ans = ans*10 + temp
	}

	if ans > max {
		return 0
	}

	return sign * ans
}
