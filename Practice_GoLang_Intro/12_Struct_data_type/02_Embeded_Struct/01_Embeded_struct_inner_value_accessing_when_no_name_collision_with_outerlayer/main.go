package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secreteAgent struct {
	person
	ltk bool
}

func main() {
	s1 := secreteAgent{
		person: person{
			first: "Rejul",
			last:  "Ramakrishnan",
			age:   32,
		},
		ltk: true,
	}
	s2 := secreteAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   35,
		},
		ltk: true,
	}
	fmt.Println(s1.first, s1.last, s1.age, s1.ltk) // s1 is able to access first name eventhough it is one level down
	fmt.Println(s2.first, s2.last, s2.age, s2.ltk) // s2.first will work as s2.person.fist
}
