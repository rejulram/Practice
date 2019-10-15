package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("nofile.txt")
	if err != nil {
		fmt.Println("Error happened", err)
	}
}
