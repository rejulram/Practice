package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

type hotdog int

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside Http")
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Method      string
		URL         *url.URL
		Host        string
		Submissions url.Values
	}{
		req.Method,
		req.URL,
		req.Host,
		req.Form,
	}

	tpl.ExecuteTemplate(res, "tpl.gohtml", data)
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)

}
