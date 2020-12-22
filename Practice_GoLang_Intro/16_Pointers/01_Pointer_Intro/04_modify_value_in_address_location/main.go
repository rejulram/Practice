package main

import "fmt"

func main(){
	a := 42
	fmt.Println("Value of a =",a)
	b := &a
	fmt.Println("Value from pointer variable *b=",*b)
	*b= 43
	fmt.Println("Value of a =",a)
}
