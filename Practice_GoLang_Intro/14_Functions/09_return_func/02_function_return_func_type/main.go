package main

import "fmt"

func main() {
	x := bar()
	fmt.Printf("Bar return type %T\n", x)
	i := x()
	fmt.Println("i := x() -->", i)

	fmt.Println("x() -->", x())

	fmt.Println("bar()() -->", bar()())
}
func bar() func() int {
	return func() int {
		return 43
	}
}
