package main

import (
	"fmt"
	"math"
)

type circle struct{
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64{ // Non pointer receiver
	return math.Pi * c.radius*c.radius
}

func info(s shape){
	fmt.Printf("%T's Area = %v",s,s.area())
}

func main(){
	c := circle{
		radius : 3.5,
	}
	info(c) //non pointer value non pointer receiver
}

