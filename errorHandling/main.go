package main

import (
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Denominator must not be zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("started Error HAndling in Functions")
	ans, _ := divide(10, 4)
	fmt.Println("Division of two number is ", ans)
}
