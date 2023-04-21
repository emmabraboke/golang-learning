package main

import "fmt"

func main() {
	addition, multiplication := addAndMultiplyNumbers(7, 5)

	fmt.Printf("7+5 = %v\n7*5 = %v", addition, multiplication)
}

// takes two integer parameter and return both the added and multiplied values
func addAndMultiplyNumbers(a int, b int) (int, int) {
	addition := a + b
	multiplication := a * b
	return addition, multiplication
}
