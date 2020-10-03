package main

import "fmt"

func main() {
	a()
	b()
	fmt.Println("Printing C ", c())
	d()
}
func a() {
	i := 0
	defer func() {
		fmt.Println("From function a", i)
	}()
	i++
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Println("From function b", i)

	}
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func d() {
	for i := 0; i < 4; i++ {
		defer func() {
			fmt.Println("From function d", i)
		}()
	}
}
