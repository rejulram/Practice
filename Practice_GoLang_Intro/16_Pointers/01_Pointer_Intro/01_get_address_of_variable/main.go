package main

import "fmt"

func main(){
	a:=42
	fmt.Println("Value of a = ",a)
	fmt.Println("Address of a =", &a)

	fmt.Printf("Type of variable a = %T\n",a)
	fmt.Printf("Type of address &a = %T\n",&a)
}
