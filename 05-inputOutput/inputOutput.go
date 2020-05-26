package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter text:")
	userText, _ := reader.ReadString('\n')
	fmt.Println(userText)

	fmt.Print("Enter number:")
	userText, _ = reader.ReadString('\n')
	userNumber, err := strconv.ParseFloat(strings.TrimSpace(userText), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(userNumber)
	}
}
