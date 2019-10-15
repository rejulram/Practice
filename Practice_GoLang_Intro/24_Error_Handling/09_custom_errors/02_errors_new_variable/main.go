package main

import (
	"errors"
	"fmt"
	"log"
)
var errorMath = errors.New("Math Error : Square root of negative number")

func main() {
	fmt.Printf("%T\n",errorMath)
	_,err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}
func sqrt(f float64) (float64,error){
	if f < 0 {
		err := errorMath
		return 0,err
	}
	return 42,nil

}