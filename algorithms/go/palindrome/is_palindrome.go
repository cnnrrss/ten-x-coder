package main

import "fmt"

func main() {
	fmt.Println("string is palindrome: ", isPalindrome("string"))
	fmt.Println("strignirts is palindrome: ", isPalindrome("stringnirts"))
}

func isPalindrome(s string) bool {
	mid := len(s) / 2
	last := len(s) - 1
	for i := 0; i < mid; i++ {
		if s[i] != s[last-i] {
			return false
		}
	}
	return true
}
