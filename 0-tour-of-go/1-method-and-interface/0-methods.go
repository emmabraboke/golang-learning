package main

import "fmt"

type sides struct {
	x, y int
}

// A method is a function with a special receiver argument
func (t sides) area() int {
	return t.x * t.y
}

func main() {
	v := sides{2, 5}

	fmt.Println(v.area())
}
