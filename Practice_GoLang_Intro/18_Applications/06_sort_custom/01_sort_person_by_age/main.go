package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Last string
	Age int
}

type byAge []Person

func (a byAge) Len() int { return len(a)}
func (a byAge) Swap(i,j int) {	a[i],a[j] = a[j],a[i]	}
func (a byAge) Less(i,j int) bool {return a[i].Age < a[j].Age }
func main(){
	p1 := Person{
		First : "James",
		Last : "Bond",
		Age : 32,
	}
	p2 := Person{
		First: "Miss",
		Last : "Moneypenny",
		Age : 28,
	}
	p3:=Person{
		"Q",
		"T",
		64,
	}
	p4:=Person{
		"W",
		"X",
		50,
	}
	People := []Person{p1,p2,p3,p4}
	fmt.Println(People)
	sort.Sort(byAge(People))
	fmt.Println(People)

}
