package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	buddha := person{
		Name:  "Buddha",
		Motto: "The belief of non beliefs",
	}
	gandhi := person{
		Name:  "Mahatma Gandhi",
		Motto: "Be the change",
	}
	mlk := person{
		Name:  "Martin Luther King ",
		Motto: "Hetred never ceases with hatred but love alone is healed",
	}
	jesus := person{
		Name:  "Jesus",
		Motto: "Love All",
	}
	muhammed := person{
		Name:  "Muhammed",
		Motto: "To overcome Evil",
	}
	persons := []person{buddha, gandhi, mlk, jesus, muhammed}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", persons)
	if err != nil {
		log.Fatalln(err)
	}
}
