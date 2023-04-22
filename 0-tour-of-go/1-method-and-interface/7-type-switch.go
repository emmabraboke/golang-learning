package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// A type switch is a construct that permits several type assertions in series.
	switch v := i.(type) {
	case int:
		fmt.Printf("Value: %v, Type: %T ", v, v)
	case string:
		fmt.Printf("Value: %v, Type: %T ", v, v)
	default:
		fmt.Printf("type '%T' not allowed", v)

	}
}
