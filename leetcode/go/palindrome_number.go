package main

import "fmt"

func isPalindrome(x int) bool {
	y := 0

	if x < 10 && x >= 0 {
		return true
	}

	if x%10 == 0 {
		return false
	}

	for x > y {
		t := x % 10
		x = x / 10
		if x == y {
			return true
		}
		y = y*10 + t
		if x == y {
			return true
		}
	}

	return false
}

func main() {
	ok := isPalindrome(10)
	fmt.Printf("Is palindrome: %v\n", ok)

	ok = isPalindrome(1)
	fmt.Printf("%v Is palindrome: %v\n", 1, ok)

	ok = isPalindrome(110)
	fmt.Printf("%v Is palindrome: %v\n", 110, ok)

	ok = isPalindrome(1101)
	fmt.Printf("%v Is palindrome: %v\n", 1101, ok)

	ok = isPalindrome(3820283)
	fmt.Printf("%v Is palindrome: %v\n", 3820283, ok)

	ok = isPalindrome(767)
	fmt.Printf("%v Is palindrome: %v\n", 7667, ok)

	ok = isPalindrome(888)
	fmt.Printf("%v Is palindrome: %v\n", 888, ok)

	ok = isPalindrome(11011)
	fmt.Printf("%v Is palindrome: %v\n", 11011, ok)
}
