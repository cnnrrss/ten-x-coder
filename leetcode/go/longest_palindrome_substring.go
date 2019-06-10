package main

import "fmt"

func longestPalindrome(s string) string {
	var resp string
	
    for i := 0; i < len(s); i++ {
		resp = expandSubstring(s, i, i, resp) // even
		resp = expandSubstring(s, i, i+1, resp) //odd
    }
    return resp
}

func expandSubstring(s string, i, j int, max string) string {
	var sub string
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i:j+1]
		i--
		j++
	}
	if len(max) < len(sub) {
		return sub
	}
	return max
}

func main() {
	fmt.Println("abccb", longestPalindrome("abcccb"))
}