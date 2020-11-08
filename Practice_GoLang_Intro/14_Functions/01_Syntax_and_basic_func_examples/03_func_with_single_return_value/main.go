package main

import "fmt"

func main() {
	s1 := woo("Moneypenny")
	fmt.Println(s1)
}

func woo(st string) string {
	return fmt.Sprintln("From Woo Hello ", st)
}
