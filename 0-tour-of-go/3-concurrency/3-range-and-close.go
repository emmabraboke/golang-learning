package main

import "fmt"

func main() {
	c := make(chan string, 3)
	go bio(c)
	for i := range c {
		fmt.Println(i)
	}
}

func bio(c chan string) {
	c <- "Hello!"
	c <- "my name is Emmanuel"
	c <- "I'm a software engineer"
	
	// A sender can close a channel to indicate that no more values will be sent
	close(c)
}
