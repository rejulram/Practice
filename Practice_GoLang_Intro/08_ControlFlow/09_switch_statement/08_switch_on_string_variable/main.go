package main

import "fmt"

func main() {
	n := "Bond"
	switch n {
	case "Rejul":
		fmt.Println("This is Rejul")
	case "Bond":
		fmt.Println("This is is Bond")
	case "Q":
		fmt.Println("This is Q")
	default:
		fmt.Println("This is default")
	}
}
