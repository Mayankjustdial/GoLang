package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "#"

func main() {
	fmt.Println("welcome to Web request")

	res, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Responnse is of type: %T\n", res)

	defer res.Body.Close() // caller responsibility to close the connection

	databytes, err := ioutil.ReadAll(res.Body)
	checkNilErr(err)

	content := string(databytes)
	fmt.Println(content)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
