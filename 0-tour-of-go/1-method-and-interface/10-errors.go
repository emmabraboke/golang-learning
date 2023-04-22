// As with Stringer, the fmt package looks for the error interface when printing values
package main

import (
	"fmt"
	"time"
)

type customError struct {
	When time.Time
	What string
}

func (e *customError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func createError() error {
	return &customError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := createError(); err != nil {
		fmt.Println(err)
	}
}
