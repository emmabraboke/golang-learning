package main

import (
	"fmt"
)

type number float64

// An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.
type Num interface {
	absolute() float64
	multiplyByTwo()
}

// You can declare a method on non-struct types, too.
func (t number) absolute() float64 {
	if t < 0 {
		return float64(-t)
	}
	return float64(t)
}

// Methods with pointer receivers can modify the value to which the receiver points
func (t *number) multiplyByTwo() {
	*t = *t * 2
}

func main() {
	var v Num
	v.multiplyByTwo()
	fmt.Println(v.absolute())
}
