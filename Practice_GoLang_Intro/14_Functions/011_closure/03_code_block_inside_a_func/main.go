package main

import "fmt"

func main() {
	var x int
	fmt.Println(x)
	x++
	fmt.Println(x)
	{
		y := 42
		fmt.Println(y)
	}
	//	fmt.Println(y)   /*Out of scope*/
	fmt.Println(x)
}

/*
Output
0
1
42
1
*/
