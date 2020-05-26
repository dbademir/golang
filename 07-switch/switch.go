package main

import "fmt"

func main() {
	var score float64

	fmt.Printf("Enter your score: ")
	fmt.Scanf("%v", &score)

	switch {
	case score <= 59:
		fmt.Println("Your grade is: F")
	case score <= 69:
		fmt.Println("Your grade is: D")
	case score <= 79:
		fmt.Println("Your grade is: C")
	case score <= 89:
		fmt.Println("Your grade is: B")
	case score <= 100:
		fmt.Println("Your grade is: A")
	default:
		fmt.Println("Please enter score lower or equal than 100")
	}
}
