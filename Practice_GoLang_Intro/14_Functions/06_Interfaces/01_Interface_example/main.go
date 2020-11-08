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

func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last)
}

type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I was passed to bar", h)
	fmt.Printf("My Type is : %T \n", h)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}
	p1 := person{
		"Dr",
		"Yes",
	}
	bar(sa1)
	bar(sa2)
	bar(p1)
}
