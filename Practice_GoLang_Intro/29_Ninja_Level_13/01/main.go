package main

import (
	"UdemyClass/Practice/29_Ninja_Level_13/01/dog"
	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
