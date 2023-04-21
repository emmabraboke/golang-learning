package main

import (
	"fmt"
	"math"
)

func main() {
	//In Go, a name is exported if it begins with a capital letter.
	// if math.ciel was use it will return error
	fmt.Println(math.Ceil(1.55))
}
