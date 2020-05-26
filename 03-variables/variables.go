package main

import (
	"fmt"
	"strconv"
)

const pi = 3.14

func main() {
	var message string
	message = "Hello, Go!"

	var message2 = "No type defined"

	var a, b, c = 1, "Go", 3
	var d float32 = 2
	message3 := "Hey, Go!"

	var myString = "3"
	myNumber, err := strconv.Atoi(myString)

	fmt.Println(myNumber)
	fmt.Println(err)

	fmt.Println(message)
	fmt.Println(message2)
	fmt.Println(a+c, b)
	fmt.Println(message3)
	fmt.Println(d * pi)
}
