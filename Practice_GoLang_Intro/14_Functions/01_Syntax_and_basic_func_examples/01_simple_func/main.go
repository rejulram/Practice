package main

import "fmt"

func main() {
	foo()
}

// FUNC Sybtax
// func (reciever)  identifier (paraameter(s)) (return(s)){code}
func foo() {
	fmt.Println("Hello from Foo")
}
