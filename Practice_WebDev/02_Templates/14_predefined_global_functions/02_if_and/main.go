package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type Person struct {
	Name  string
	Motto string
	Admin bool
}

func main() {
	p1 := Person{
		Name:  "Gandhi",
		Motto: "Be the change",
		Admin: true,
	}
	p2 := Person{
		Name:  "Buddha",
		Motto: "Belief of no belief",
		Admin: false,
	}
	p3 := Person{
		Name:  "",
		Motto: "Nobody",
		Admin: true,
	}
	persons := []Person{p1, p2, p3}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", persons)
	if err != nil {
		log.Fatalln(err)
	}
}
