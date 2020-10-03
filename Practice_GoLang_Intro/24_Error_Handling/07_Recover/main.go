package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned Normally from f.")
}
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned Normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicing !!")
		panic(fmt.Sprintf("%v", i*2))
	}
	defer fmt.Println("Defer in g.", i)
	fmt.Println("Printing in g.", i)
	g(i + 1)
}
