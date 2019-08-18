package main

import "fmt"

// Write a function to swap two numbers in place (that is, without temporary variables).
func main() {
	a, b := 9, 5
	fmt.Println("initial:", a, b)
	// addition method
	a = a - b
	b = b + a
	a = b - a
	fmt.Println("swapped:", a, b)

	// XOR method
	a ^= b
	b ^= a
	a ^= b

	fmt.Println("swap again:", a, b)
}
