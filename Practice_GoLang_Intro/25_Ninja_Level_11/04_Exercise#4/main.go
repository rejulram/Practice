package main

import (
	"fmt"
)

type sqrtError struct {
	lat string
	long string
	err error
}

func (s sqrtError) Error() string {
	return fmt.Sprintf("Math Error : %v %v %v",s.lat,s.long,s.err)
}
func main() {
	_,err := sqrt(-10.23)
	if err != nil {
		fmt.Println(err)
	}
}
func sqrt(f float64) (float64,error) {
	if f < 0{
		//e := errors.New("negative number found")
		e := fmt.Errorf("negative number found %v",f)
		return 0,sqrtError{"15.0045","45.4534",e}
	}
	return 42,nil
}
