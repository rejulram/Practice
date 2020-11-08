package main

import "fmt"

var x, y int //Multiple variable declaration

func main() {
	x, y = 43, 45 //Multiple variable assignement
	x, y = y, x   // Swap the numbers
	fmt.Println(x, y)
}
