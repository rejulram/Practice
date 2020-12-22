package main

import "fmt"

func main() {
	func(x int) {
		fmt.Println("Anonymous func with parameter ran :", x)
	}(42)
}
