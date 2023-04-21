package main

import "fmt"

func main() {
	addition, multiplication := addAndMultiplyNumbers(7, 5)

	fmt.Printf("7+5 = %v\n7*5 = %v", addition, multiplication)
}

// takes two integer parameter and return both the added and multiplied values
// This func uses a named return value
// A return statement without arguments returns the named return values. This is known as a "naked" return.
func addAndMultiplyNumbers(a int, b int) (addition, multiplication int) {
	addition = a + b
	multiplication = a * b
	return
}
