package main

import "fmt"

func main() {
	func(x int) {
		fmt.Println("Anonymous func with paramter ran :", x)
	}(42)
}
