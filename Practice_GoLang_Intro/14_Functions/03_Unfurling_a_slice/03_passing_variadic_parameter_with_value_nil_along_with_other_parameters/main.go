package main

import "fmt"

func main() {
	output := sum("Rejul") // No values passed
	fmt.Println("Sum is", output)
}

func sum(s string, x ...int) int {
	fmt.Println("I am ", s)
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("For item in index position", i, "we are adding", v, "to the total", sum)
	}
	return sum
}
