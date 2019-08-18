package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Create a stack-based evaluator for an expression in reverse Polish notation (RPN).
// that also shows the changes in the stack as each individual token is processed as a table.

func main() {
	input := "3 4 2 * 1 5 - 2 3 ^ ^ / +"
	result := calculate(input)
	fmt.Println("\nThe final value is", result)
	input = "3 4 * 2 / 1"
	result = calculate(input)
	fmt.Println("\nThe final value is", result)
}

func calculate(input string) float64 {
	fmt.Println("\nToken            Action            Stack")
	stack := []float64{}
	for _, tok := range strings.Fields(input) {
		action := "Apply op to top of stack"
		switch tok {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/":
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "^":
			stack[len(stack)-2] =
				math.Pow(stack[len(stack)-2], stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		default:
			action = "Push num onto top of stack"
			f, _ := strconv.ParseFloat(tok, 64)
			stack = append(stack, f)
		}
		fmt.Printf("%3s    %-26s  %v\n", tok, action, stack)
	}
	return stack[0]
}
