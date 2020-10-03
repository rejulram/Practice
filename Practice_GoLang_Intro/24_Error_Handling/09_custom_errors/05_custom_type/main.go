package main

import (
	"fmt"
	"log"
)

type mathError struct {
	lat string
	lon string
	err error
}

func (m mathError) Error() string {
	return fmt.Sprintf("Math Error occured : %v,%v,%v", m.lat, m.lon, m.err)
}

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 10 {
		err := fmt.Errorf("sqauare root of Negative number is invalid  %v", f)
		return 0, mathError{"10.4445 N", "45.5543 W", err}
	}
	return 42, nil
}
