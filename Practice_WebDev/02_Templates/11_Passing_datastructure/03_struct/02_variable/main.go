package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name string
	Motto string
}

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p1 := person {
		Name : "Buddha",
		Motto : "Belief of non belief",
	}
	err := tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",p1)
	if err != nil {
		log.Fatalln(err)
	}
}