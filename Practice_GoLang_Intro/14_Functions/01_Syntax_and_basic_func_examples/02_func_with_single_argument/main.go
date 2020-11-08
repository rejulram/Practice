package main

import "fmt"

func main() {
	bar("James")
}

//Everything in GO is Pass by value
func bar(s string) {
	fmt.Println("Hello", s)
}
