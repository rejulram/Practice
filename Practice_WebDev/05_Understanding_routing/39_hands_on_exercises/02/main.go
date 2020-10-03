package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("something.gohtml"))
}

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome main page")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to dog's Page")
}

func me(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "something.gohtml", "Rejul")
	if err != nil {
		log.Fatalln("Execute template failed")
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
