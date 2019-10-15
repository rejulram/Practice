package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println("Cannot create file :", err)
		return
	}
	defer f.Close()
	r := strings.NewReader("Hello Rejul")
	io.Copy(f, r)
}
