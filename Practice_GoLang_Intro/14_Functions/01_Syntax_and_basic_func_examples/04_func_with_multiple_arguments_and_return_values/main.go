package main

import "fmt"

func main() {
	x, y := mouse("Rejul", "Ramakrishnan")
	fmt.Println(x)
	fmt.Println(y)
}

func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `! says "Hello"`)
	b := false
	return a, b
}
