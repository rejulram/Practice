package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Rejul Ramakrishnan"
	tpl := `
	<!DOCTYPE html>
	<html lang="en"">
    <head>
     <meta charset="UTF-8">
     <title>Hello world</title>
	</head>
	 <body>
     <h1>` + name + `</h1>
     </body>
	</html>
`
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(tpl))
}
