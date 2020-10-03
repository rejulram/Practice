package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}
func sqrt(f float64) (float64, error) {
	if f < 0 {
		errMath := fmt.Errorf("Math Error: Trying take sqaure root of neagive number %v", f)
		return 0, errMath
	}
	return 42, nil

}
