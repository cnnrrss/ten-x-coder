package main

import "fmt"

func reverse(x int) int {
	var sign, temp, ans int
	var max int = 1<<31 - 1

	if x > 0 {
		sign = 1
	} else {
		sign = -1
		x = -x
	}

	for x >= 10 {
		temp = x % 10
		x = x / 10
		ans = ans*10 + temp
	}

	temp = x % 10

	if temp > 0 {
		ans = ans*10 + temp
	}

	if ans > max {
		return 0
	}

	return ans * sign
}

func main() {
	ans := reverse(1234)
	fmt.Printf("%v reversed = %v\n", 1234, ans)
	ans = reverse(-1234)
	fmt.Printf("%v reversed = %v\n", -1234, ans)
	ans = reverse(210)
	fmt.Printf("%v reversed = %v\n", 210, ans)
	ans = reverse(2147483648)
	fmt.Printf("%v reversed = %v\n", 2147483648, ans)
	ans = reverse(10)
	fmt.Printf("%v reversed = %v\n", 10, ans)
	fmt.Printf("%d | %d = %d", 10, 10, 10%10)
}
