package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string)
	go greet(c)
	go reply(c)

	time.Sleep(100 * time.Millisecond)
	reply := <-c
	fmt.Println(reply)
}

func greet(message chan string) {
	message <- "Hello!"
}

func reply(message chan string) {
	msg := <-message

	fmt.Println(msg)
	time.Sleep(1000 * time.Millisecond)

	message <- "Hi"
}
