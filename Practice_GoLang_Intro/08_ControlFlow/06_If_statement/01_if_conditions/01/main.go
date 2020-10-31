package main

import "fmt"

func main(){
	if true{
		fmt.Println("Will print 001")
	}
	if false{
		fmt.Println("Won't print 002")
	}
	if !true{
		fmt.Println("Won't print 003")
	}
	if !false{
		fmt.Println("Will print 004")
	}
	if 2==2 {
		fmt.Println("Will print 005")
	}
	if 2!=2 {
		fmt.Println("Won't print 006")
	}
}
