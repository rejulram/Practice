package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "Rejul",
		last:  "Ramakrishnan",
	}
	p2 := person{
		first: "James",
		last:  "Bond",
	}
	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)
}
