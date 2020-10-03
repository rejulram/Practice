package main

import "fmt"

var a int8 = -128 // will work
//var a int8 = -129 // won't work
//var a int8 = 127 // will work
//var a int8 = 128 // won't work

func main() {
	fmt.Println("a =", a)
	fmt.Printf("Type of a is %T\n", a)
}
