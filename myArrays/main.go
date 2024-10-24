package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitList [4]string
	fruitList[0] = "Mango"
	fruitList[1] = "Apple"
	fruitList[3] = "Papaya"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [3]string{"patato", "tamato", "mushroom"}
	fmt.Println("Vegy List is: ", vegList)
	fmt.Println("Vegy List is: ", len(vegList))

}
