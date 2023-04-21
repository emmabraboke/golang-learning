package main

import "fmt"

func main() {
	printName("Emmanuel", "Braboke")
	name := fullName("Emmanuel", "Braboke")
	fmt.Println("my name is", name)
}

//This function receives to string parameter and prints the concatenated strings
func printName(firstName string, lastName string) {
	fmt.Println(lastName + " " + lastName)
}

//This function receives to string parameter and returns the concatenated strings
func fullName(firstName string, lastName string) string {
	return lastName + " " + lastName
}
