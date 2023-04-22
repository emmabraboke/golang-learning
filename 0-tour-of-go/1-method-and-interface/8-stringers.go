// A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values
package main

import "fmt"

type name struct {
	firstName, lastName string
}

func (t name) String() string {
	return fmt.Sprintf("my name is %v %v", t.firstName, t.lastName)
}

func main() {
	fullName := name{"Emmanuel", "Braboke"}
    
	fmt.Println(fullName)
}
