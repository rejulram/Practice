package main

import "fmt"

var a int
var b float64

func main() {
	a = 10
	b = 10.1234
	// a = 10.1234 // won't run
	fmt.Println("a =", a)
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("b =", b)
	fmt.Printf("Type of b is %T\n", b)
}
