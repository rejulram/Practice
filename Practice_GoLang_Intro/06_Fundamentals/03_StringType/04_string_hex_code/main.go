package main

import "fmt"

func main() {
	s := "ഞാൻ രജുൽ"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}

	for i, v := range s {
		fmt.Printf("Index position %d we have value hex %x\n", i, v)
	}
}
