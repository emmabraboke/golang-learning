package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	var list List[int]
	var nextList List[int]
	list.val = 1 
	nextList.val = 2
	nextList.next = &list

	fmt.Println(nextList)
}
