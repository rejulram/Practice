package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template
type hotdog int

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func (h hotdog) ServeHTTP(w http.ResponseWriter,req *http.Request){
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	data := struct{
		Method string
		Submissions url.Values
	}{
		req.Method,
		req.Form,
	}
	tpl.ExecuteTemplate(w,"tpl.gohtml",data)
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080",d)
}

