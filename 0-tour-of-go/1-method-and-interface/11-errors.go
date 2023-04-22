package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (t ErrNegativeSqrt) Error() string {
	err := fmt.Sprintf("cannot Sqrt negative number: %v", float64(t))
	return err
}
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	//square root
    z := x / 2
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(9))
	fmt.Println(Sqrt(-9))
}
