package main

import "fmt"

func PrintValues[T any](x, y T) {
	fmt.Println(x,y)
}

func main(){
	PrintValues("Hello","Emmanuel")
	PrintValues(22,2023)
}