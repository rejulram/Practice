package main

import "fmt"

func main() {
	x := func() int {
		return 42
	}
	fmt.Printf("%T\n", x)
}
