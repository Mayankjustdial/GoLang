package main

import "fmt"

func main() {
	fmt.Println("Structs in goLang")
	// no inheritance in golang; No super or parent

	mayank := User{"Mayank", "mayank@go.dev", true, 24}
	fmt.Println(mayank)
	fmt.Printf("Mayank details are: %+v\n", mayank)
	fmt.Printf("Name ia %v and email is %v.\n", mayank.Name, mayank.Email)
	mayank.GetStatus()
	mayank.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", mayank.Name, mayank.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user acive: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("email of this uset is: ", u.Email)
}
