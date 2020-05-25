package main

import "fmt"

const pi = 3.14

func main() {
	var message string
	message = "Hello, Go!"

	var message2 = "No type defined"

	var a, b, c = 1, "Go", 3
	var d float32 = 2
	message3 := "Hey, Go!"

	fmt.Println(message)
	fmt.Println(message2)
	fmt.Println(a+c, b)
	fmt.Println(message3)
	fmt.Println(d * pi)
}
