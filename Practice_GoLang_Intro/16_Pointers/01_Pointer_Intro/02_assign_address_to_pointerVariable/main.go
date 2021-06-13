package main

import "fmt"

func main(){
	a := 42
	fmt.Println("Value of a =",a)
	fmt.Println("Address of &a =",&a)
	var b *int = &a
	fmt.Println("Value of b =",b)
	fmt.Printf("Type of b =%T\n",b)
}
