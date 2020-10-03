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
func main() {
	person := map[string]string{
		"India":   "Gandhi",
		"USA":     "MLK",
		"Love":    "Jesus",
		"Prophet": "Mohammed",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", person)
	if err != nil {
		log.Fatalln(err)
	}
}
