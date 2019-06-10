package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	bracketsMaps := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, v := range s {
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}
		// fmt.Println(i, string(stack[len(stack)-1]), string(bracketsMaps[v]), string(v))
		if bracketsMaps[v] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
		// Could be greedier
	}
	return len(stack) == 0
}

func main() {
	ok := isValid("{}")
	fmt.Println(ok)

	ok = isValid("()")
	fmt.Println(ok)
	ok = isValid("([{[]}])")
	fmt.Println(ok)
	ok = isValid("[[[[[](]]]]")
	fmt.Println(ok)
	ok = isValid("}{")
	// ok = isValid("(]")
	fmt.Println(ok)
}
