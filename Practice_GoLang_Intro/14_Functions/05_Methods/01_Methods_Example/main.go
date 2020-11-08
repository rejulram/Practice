package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r reciever) identifiers(parameters(s)) (return(s)) {code.. }

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}
func main() {
	s1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	s2 := secretAgent{
		person: person{
			"Miss",
			"MoneyPenny",
		},
		ltk: true,
	}
	s1.speak()
	s2.speak()
}
