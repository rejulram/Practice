package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
type Person struct {
	Name  string
	Motto string
}

type Car struct{
	Manufacture  string
	Model string
	Doors int
}

type Items struct{
	Wisdom []Person
	Transport []Car
}
func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	buddha := Person {
		Name : "Buddha",
		Motto : "The belief of non beliefs",
	}
	gandhi := Person{
		Name : "Mahatma Gandhi",
		Motto : "Be the change",
	}
	mlk := Person{
		Name : "Martin Luther King ",
		Motto : "Hetred never ceases with hatred but love alone is healed",
	}
	jesus := Person {
		Name: "Jesus",
		Motto: "Love All",
	}
	muhammed := Person {
		Name : "Muhammed",
		Motto : "To overcome Evil",
	}
	ford := Car {
		Manufacture: "Ford",
		Model: "Figo",
		Doors: 4,
	}
	toyota := Car{
		Manufacture: "Toyota",
		Model : "Corolla",
		Doors:4,
	}
	persons := []Person{buddha,gandhi,mlk,jesus,muhammed}
	cars := []Car{ford,toyota,}
	data := Items{
		Wisdom: persons,
		Transport: cars,
	}
	err := tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",data)
	if err != nil {
		log.Fatalln(err)
	}
}
