// There are two reasons to use a pointer receiver.
// The first is so that the method can modify the value that its receiver points to.
// The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct
package main

import (
	"fmt"
)

type number float64

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
	v := number(-1)
	v.multiplyByTwo()
	fmt.Println(v.absolute())
}
