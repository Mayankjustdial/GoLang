package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the pizza: ")

	// comma ok || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T, ", input)

}
