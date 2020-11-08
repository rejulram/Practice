package main

import "fmt"

func main() {
	s := []int{2, 3, 4, 5, 6, 7}
	output := sum(s...)
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
