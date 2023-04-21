package main

import "fmt"

//constant value cannot be reassigned after declaration
const g = 9.81

const (
	// shifting a 1 bit left 10 places
	big = 1 << 10
	// Shift it right again 9 places, so we end up with 1<<1, or 2.
	small = big >> 9
)

func main() {
	fmt.Printf("acceleration due to gravity, g is %v m/s2\n", g)

	fmt.Println(big)
	fmt.Println(small)
}
