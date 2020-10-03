package main

import (
	"errors"
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
		err := errors.New("Math Error : Square root of negative number")
		return 0, err
	}
	return 42, nil

}
