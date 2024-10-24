package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to topic on Slices ")

	var fruitList = []string{"Apple", "Tomato", "Mango"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Banana", "Grapes")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScore := make([]int, 5)
	highScore[0] = 234
	highScore[1] = 345
	highScore[2] = 456
	highScore[3] = 123
	highScore[4] = 143
	// highScore[4] = 453

	highScore = append(highScore, 222, 452, 444)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

}
