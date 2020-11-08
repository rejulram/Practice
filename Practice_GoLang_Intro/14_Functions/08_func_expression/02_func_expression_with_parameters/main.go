package main

import "fmt"

func main() {
	g := func(x int) {
		fmt.Println("Func expression with int argument", x)
	}
	g(42)
}
