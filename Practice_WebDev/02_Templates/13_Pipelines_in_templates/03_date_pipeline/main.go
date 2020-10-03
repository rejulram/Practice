package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fdate": formatDate,
}

func init() {
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func formatDate(t time.Time) string {
	return t.Format("02-01-2006")
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
