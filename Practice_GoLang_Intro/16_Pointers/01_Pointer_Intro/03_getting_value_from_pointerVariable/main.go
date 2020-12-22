package main

import "fmt"

func main(){
	a := 42
	fmt.Println("Value of a =",a)
	fmt.Printf("Address of &a =%T\n",&a)
	var b *int = &a
	fmt.Println("Value of b =",b)
	fmt.Printf("Type of a =%T\n",b)

	fmt.Println("Value from pointer variable *b=",*b)
	fmt.Println("Value from a variable using pointer symbol",*&a)
}
