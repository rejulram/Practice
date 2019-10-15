package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("nofile.txt")
	if err != nil {
		log.Panic(err)
	}
}
func foo() {
	fmt.Println("This defer function will run")
}
