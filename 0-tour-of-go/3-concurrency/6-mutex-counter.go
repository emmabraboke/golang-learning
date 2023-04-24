package main

import (
	"fmt"
	"sync"
)

var mu = sync.Mutex{}
var wg = sync.WaitGroup{}
var counter = 0


func main() {
	
	for i := 0; i < 10; i++ {
		wg.Add(2)
		mu.Lock()
		go sayHello()
		mu.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("hello %v\n", counter)
	mu.Unlock()
	wg.Done()
}

func increment() {
	counter++
	mu.Unlock()
	wg.Done()
}