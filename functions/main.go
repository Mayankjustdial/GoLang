package main

import "fmt"

func main() {
	fmt.Println("welcome to function in goLang")
	greater()

	greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result is:", result)

	proResult, myMessage := proAdder(3, 46, 2, 23, 5, 6)
	fmt.Println("Pro result is : ", proResult)
	fmt.Println("Pro messsage is : ", myMessage)

}

func adder(valOne int, valtwo int) int {
	return valOne + valtwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi pro result function"
}

func greater() {
	fmt.Println("Namaste form goLang")
}

func greeterTwo() {
	fmt.Println("Another method")
}
