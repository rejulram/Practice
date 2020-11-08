package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secreteAgent struct {
	person
	first string
	ltk   bool
}

func main() {
	s1 := secreteAgent{
		person: person{
			first: "Rejul",
			last:  "Ramakrishnan",
			age:   32,
		},
		first: "Name Collision",
		ltk:   true,
	}
	s2 := secreteAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   35,
		},
		ltk: true,
	}
	fmt.Println(s1.person.first, s1.last, s1.age, s1.ltk) // s1 is not able to access first name as there is name collission with first
	fmt.Println(s2.person.first, s2.last, s2.age, s2.ltk) //
}
