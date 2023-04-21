// Every go program starts running in package main
package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("print a random number", rand.Float32())
	fmt.Println("The absolute value of -100 is", math.Abs(-100))
}
