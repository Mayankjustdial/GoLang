package main

import "fmt"

func main() {
	fmt.Println("Structs in goLang")
	// no inheritance in golang; No super or parent

	mayank := User{"Mayank", "mayank@go.dev", true, 24}
	fmt.Println(mayank)
	fmt.Printf("Mayank details are: %+v\n", mayank)
	fmt.Printf("Name ia %v and email is %v.", mayank.Name, mayank.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
