package main

import "fmt"

func main() {
	fmt.Println("5+6 =",mySum(5,6))
	fmt.Println("7+5 =",mySum(7,5))
	fmt.Println("9+4 =",mySum(9,4))

}
func mySum(xi ...int) int{
	sum:=0
	for _,v := range xi {
		sum += v
	}
	return sum
}
