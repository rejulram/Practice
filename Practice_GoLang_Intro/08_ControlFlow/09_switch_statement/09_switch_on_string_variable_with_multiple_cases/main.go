package main

import "fmt"

func main(){
	n := "Bond"
	switch n {
	case "Rejul","Bond","James" :
		fmt.Println("This is Rejul or Bond or James")
	case "M" :
		fmt.Println("This is is M")
	case "Q" :
		fmt.Println("This is Q")
	default:
		fmt.Println("This is default")
	}
}

