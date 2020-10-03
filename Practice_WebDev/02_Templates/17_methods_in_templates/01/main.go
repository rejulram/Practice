package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type Person struct {
	Name string
	Age  int
}

func (p Person) Something() int {
	return 7
}

func (p Person) DblAge() int {
	return p.Age * 2
}

func (p Person) TakeArg(x int) int {
	return x * 2

}
func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	p1 := Person{
		Name: "Rejul",
		Age:  33,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", p1)
	if err != nil {
		log.Fatalln(err)
	}
}
