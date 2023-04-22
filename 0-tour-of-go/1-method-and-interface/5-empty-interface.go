package main

import "fmt"

func main() {
  var i interface{}

  i = "hi"
  describe(i)

  i = 1

  describe(i)
}

func describe(i interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", i, i)
}