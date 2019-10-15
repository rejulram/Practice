package main

import (
	"fmt"
	"log"
)

func main() {
	_,err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}
func sqrt(f float64) (float64,error){
	if f < 0 {
		return 0,fmt.Errorf("Math Error: Trying take sqaure root of neagive number %v",f)
	}
	return 42,nil

}
