package main

import "fmt"

type Name struct {
	firstName, lastName string
}

type Names interface {
	printName()
}

// This method means type Name implements the interface Names,
// but we don't need to explicitly declare that it does so.
func (t Name) printName() {
	fmt.Println(t.firstName + " " + t.lastName)
}

func main() {
	var name Names = Name{"Emmanuel", "braboke"}
	name.printName()
}
