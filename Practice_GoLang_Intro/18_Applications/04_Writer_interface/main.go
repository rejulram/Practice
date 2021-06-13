package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	fmt.Println("Hello Rejul")
	fmt.Fprintln(os.Stdout,"Hello Rejul")
	io.WriteString(os.Stdout,"Hello Rejul")
}