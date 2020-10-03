package main

import (
	"fmt"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("nofile.txt")
	if err != nil {
		panic(err)
	}
}
func foo() {
	fmt.Println("This defer function will run")
}
