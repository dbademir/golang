package main

import (
	"fmt"
	"os"
)

func main() {
	uHome := os.Getenv("HOME")
	uGoRoot := os.Getenv("GOROOT")
	uGoPath := os.Getenv("GOPATH")
	fmt.Println("Home:" + uHome)
	fmt.Println("Go Root:" + uGoRoot)
	fmt.Println("Go Path:" + uGoPath)

}

/*
func main() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}
*/
