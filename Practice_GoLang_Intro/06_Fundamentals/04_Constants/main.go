package main

import "fmt"

const a = 42
const b = 42.234
const c = "Hello"

/*


const (
	a = 42
	b = 42.2345
	c = "Hello"
)

const (
	a int = 42
	b float64 = 42.2345
	c string = "Hello"
)



*/

func main() {
	fmt.Println("a =", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Printf("Type of a : %T\n", a)
	fmt.Printf("Type of b : %T\n", b)
	fmt.Printf("Type of c : %T\n", c)
}
