package main

import "fmt"

type customErr struct {
	info string
}

func (c customErr) Error() string {
	return fmt.Sprintln("There is an error : ", c.info)

}

func main() {
	e := customErr{
		info: "Need more cofee",
	}
	foo(e)
}
func foo(e error) {
	//fmt.Println("foo ran -",e)
	//Following code won't work
	//fmt.Println("foo ran -",e.info)
	// Need do assertion then above code works
	fmt.Println("foo ran -", e.(customErr).info)
}
