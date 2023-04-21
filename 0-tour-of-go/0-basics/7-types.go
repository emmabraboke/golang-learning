package main

import "fmt"

func main() {

	// complex number
	a := 3 + 7i
	b := complex(3, 7)
	c := 1 << 10

	//boolen
	isTrue := true

	fmt.Printf("Type: %T Value: %v\n", a, a)
	fmt.Printf("Type: %T Value: %v\n", b, b)
	fmt.Printf("Type: %T Value: %v\n", c, c)
	fmt.Printf("Type: %T Value: %v\n", isTrue, isTrue)
}
