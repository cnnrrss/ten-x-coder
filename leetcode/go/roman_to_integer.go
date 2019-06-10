package main

import "fmt"

func romanToInt(s string) int {
	sum := 0
	carry := 0
	convert := map[rune]int{
		73: 1,
		86: 5,
		88: 10,
		76: 50,
		67: 100,
		68: 500,
		77: 1000,
	}
	for i, v := range s {
		if carry < convert[v] {
			fmt.Printf("Loop:%v\nSum:%v\nWiz:%v\nDict:%v\n", i, sum, carry*2, convert[v])
			fmt.Println("#######################")
			sum = sum - (carry * 2) + convert[v]
		} else {
			sum = sum + convert[v]
		}
		carry = convert[v]
	}
	return sum
}

func main() {
	// s := romanToInt("IVXLCDM")
	// fmt.Println(s)
	s := romanToInt("MCMXCIV")
	fmt.Println(s)
}
