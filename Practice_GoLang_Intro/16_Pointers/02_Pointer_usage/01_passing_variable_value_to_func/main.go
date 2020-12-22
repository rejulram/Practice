package main

import "fmt"

func main(){
	x:=42
	foo(x)
	fmt.Println("After x=",x)
}

func foo(y int){
	fmt.Println("before y=",y)
	y = 43
	fmt.Println("After y=",y)
}
