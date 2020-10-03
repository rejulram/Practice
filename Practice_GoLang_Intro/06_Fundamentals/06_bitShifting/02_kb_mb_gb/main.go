package main

import "fmt"

func main() {
	kb := 1024
	mb := kb * 1024
	gb := mb * 1024
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
