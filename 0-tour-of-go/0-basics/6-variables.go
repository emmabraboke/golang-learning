package main

import "fmt"

var firstName, lastName string

func main() {
	//variable with initializer
	var profession, language = "software engineer", "golang"

	// short variable declaration
	symbol, newLanguage := "!", "node.js"
	firstName = "Emmanuel"
	lastName = "Braboke"

	fmt.Println("My name is", lastName, firstName, "I'm a", profession, "with experience using", language)
	fmt.Println("I also use", newLanguage+symbol)
}
