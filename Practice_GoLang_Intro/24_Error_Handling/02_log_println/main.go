package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("nofile.txt")
	if err != nil {
		log.Println("Error occurred", err)
	}
}
