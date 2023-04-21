package main

import "fmt"

func main() {
	var a int = 3
	var b float32 = 6.25

	fmt.Printf("Value: %v Type: %T\n", int(b), int(b))
	fmt.Printf("Value: %v Type: %T\n", float32(a), float32(a))
}
