package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	bracketsMaps := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for i, v := range s {
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}
		fmt.Println(i, stack[len(stack)-1], bracketsMaps[v], v)
		if bracketsMaps[v] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack) == 0
}

func main() {
	ok := isValid("{}")
	fmt.Println(ok)
}
