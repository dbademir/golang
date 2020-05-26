package main

import (
	"fmt"
	"time"
)

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

	for i := 0; i <= 9; i++ {
		fmt.Println("i: ", i)
	}

	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
		time.Sleep(500 * time.Millisecond)
	}
}
