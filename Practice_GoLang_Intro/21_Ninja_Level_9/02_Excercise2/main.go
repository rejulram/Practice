package main

import "fmt"

type person struct{
	Name string
}

func (p *person) speak(){
	fmt.Println("Hi I am ",p.Name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{"Rejul"}
	saySomething(&p)
}