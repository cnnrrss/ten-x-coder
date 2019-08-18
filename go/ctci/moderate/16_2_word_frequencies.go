package main

import (
	"fmt"
	"strings"
)

func main() {
	book := []string{"The", "big", "bad", "wolf", "went", "walking", "in", "the", "big", "bad", "woods"}
	fmt.Println(countFrequenciesSimple(book, "bad"))
	fmt.Println(countFrequenciesRepetitive(book, "bad"))
	fmt.Println(countFrequenciesSimple(book, "wolf"))
	fmt.Println(countFrequenciesRepetitive(book, "wolf"))
}

func countFrequenciesSimple(book []string, word string) int {
	count := 0
	for _, str := range book {
		if str == word {
			count++
		}
	}
	return count
}

func countFrequenciesRepetitive(book []string, word string) int {
	dict := setupDictionary(book)
	return getFrequency(dict, word)
}

func setupDictionary(book []string) map[string]int {
	dict := map[string]int{}
	for _, word := range book {
		strings.ToLower(word)
		dict[word]++
	}
	return dict
}

func getFrequency(dict map[string]int, word string) int {
	if _, ok := dict[word]; ok {
		return dict[word]
	}
	return 0
}
