package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// fmt.Println("Hey, What's your name?")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Hello, Mr. ", name)

	welcome := "welcome to user Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the pizza: ")

	// comma ok || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T, ", input)
}
