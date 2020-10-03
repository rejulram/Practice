package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type Person struct {
	Name  string
	Motto string
}

var fm = template.FuncMap{
	"up":    strings.ToUpper,
	"strip": firstThree,
}

func init() {
	tpl = template.Must(template.New("myTemplate").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	buddha := Person{
		Name:  "Buddha",
		Motto: "The belief of non beliefs",
	}
	gandhi := Person{
		Name:  "Mahatma Gandhi",
		Motto: "Be the change",
	}
	mlk := Person{
		Name:  "Martin Luther King ",
		Motto: "Hatred never ceases with hatred but love alone is healed",
	}
	jesus := Person{
		Name:  "Jesus",
		Motto: "Love All",
	}
	muhammed := Person{
		Name:  "Muhammed",
		Motto: "To overcome Evil",
	}
	persons := []Person{buddha, gandhi, mlk, jesus, muhammed}
	data := struct {
		Wisdom []Person
	}{
		Wisdom: persons,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
