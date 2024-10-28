package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all Language: ", languages)
	fmt.Println("JS short for: ", languages["JS"])

	// delete(languages, "RB")
	fmt.Println("List of all Language: ", languages)

	// loops are interesting in GoLang

	for key, value := range languages {
		fmt.Printf("For key %v,value is %v\n ", key, value)
	}

	for _, value := range languages {
		fmt.Printf("For key v,value is %v\n ", value)
	}
}
