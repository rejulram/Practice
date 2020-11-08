package main

import "fmt"

func main() {
	fmt.Println(4 * 3 * 2 * 1)
	f := factorial(4)
	fmt.Println(f)

	f2 := factLoop(4)
	fmt.Println(f2)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factLoop(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
