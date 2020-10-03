package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	c1 := circle{5}
	// info(&c1) // works
	info(c1) // Won't work
}
