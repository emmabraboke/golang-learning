package main

import (
	"fmt"
)

type number float64

// You can declare a method on non-struct types, too.
func(t number) absolute() float64{
	if t < 0 {
		return float64(-t)
	}
	return float64(t)
}

func main() {
	v := number(-1) 

	fmt.Println(v.absolute())
}
