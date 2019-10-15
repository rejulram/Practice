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
		log.Fatalln(err)
	}
}
func foo() {
	fmt.Println("This defer function won't run")
}
