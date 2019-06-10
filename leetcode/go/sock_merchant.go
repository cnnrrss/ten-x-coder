package main

import "fmt"

func sockMerchant(n int32, ar []int32) int32 {
	inventory := make(map[int32]int32)
	var pairs int32
	for _, i := range ar {
		if inventory[i] != 0 {
			pairs++
			inventory[i] = 0
		} else {
			inventory[i]++
		}
	}
	return pairs
}

func main() {
	fmt.Println(sockMerchant(4, []int32{1, 1, 3, 3, 4, 4}))
}
