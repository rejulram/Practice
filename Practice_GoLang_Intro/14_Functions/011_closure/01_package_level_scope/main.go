package main

import "fmt"

var x int

func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
	fmt.Println(x)
}

func foo() {
	fmt.Println("Hello foo")
	x++
}

/*
Output
0
1
Hello foo
2
*/
