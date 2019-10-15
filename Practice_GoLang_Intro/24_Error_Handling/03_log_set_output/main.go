package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f1, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Error happened while Creating file", err)
	}
	defer f1.Close()
	log.SetOutput(f1)
	f2, err := os.Open("nofile.txt")
	if err != nil {
		log.Println("Error happened file reading file", err)
	}
	defer f2.Close()
	fmt.Println("Please check log.txt file in the directory")
}
