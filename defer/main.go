package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")

	myDefer()
}

// World, One, Two
// 0,1,2,3,4
// Output: Hello 43210 Two One World

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
