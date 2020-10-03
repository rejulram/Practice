package main

import "fmt"

func main() {
	s := "Hello Rejul"
	fmt.Println(s)
	fmt.Printf("Type of s is %T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("Type of bs is %T\n", bs)
}
