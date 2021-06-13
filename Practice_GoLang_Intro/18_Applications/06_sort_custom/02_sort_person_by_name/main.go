package main

import (
	"fmt"
	"sort"
)

type person struct {
	First  string
	Age int
}

type byName []person

func (bs byName) Len() int { return len(bs)}
func (bs byName) Swap(i,j int) { bs[i],bs[j] = bs[j],bs[i] }
func (bs byName) Less(i,j int) bool { return bs[i].First < bs[j].First }

func main(){
	p1 := person{"James",32,}
	p2 := person{"Moneypenny",28}
	p3 := person{"Q",54}
	p4 := person{"M",68}
	people := []person{p1,p2,p3,p4}
	fmt.Println(people)
	sort.Sort(byName(people))
	fmt.Println(people)
}