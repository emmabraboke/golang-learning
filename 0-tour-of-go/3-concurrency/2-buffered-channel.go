package main

import "fmt"

func main() {
	c := make(chan string, 3)

	c <- "Hello!"
	c <- "my name is Emmanuel"
	c <- "I'm a software engineer"

	greeting := <-c
	name := <-c
	profession := <-c

	fmt.Printf("%v, %v, %v.", greeting, name, profession)

}

func bio(message chan string){
	
}