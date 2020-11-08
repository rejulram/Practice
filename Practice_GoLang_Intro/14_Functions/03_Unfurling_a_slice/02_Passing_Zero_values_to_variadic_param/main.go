package main

import "fmt"

func main() {
	output := sum() // No values passed
	fmt.Println("Sum is", output)
}

func sum(x ...int) int {
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("For item in index position", i, "we are adding", v, "to the total", sum)
	}
	return sum
}
