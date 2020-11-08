package main

import "fmt"

func main() {
	s := []int{2, 3, 4, 5, 6, 7}
	output := sum("Rejul", s...)
	fmt.Println("Sum is", output)
}

func sum(s string, x ...int) int { // Variadic parameter should be final parameter
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("For item in index position", i, "we are adding", v, "to the total", sum)
	}
	return sum
}
