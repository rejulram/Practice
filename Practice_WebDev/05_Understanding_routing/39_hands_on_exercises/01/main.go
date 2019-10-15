package main

import (
	"io"
	"net/http"
)

func index(res http.ResponseWriter,req *http.Request){
	io.WriteString(res,"Welcome Index page")
}

func dog(res http.ResponseWriter,req *http.Request){
	io.WriteString(res,"Welcome Dog page")
}

func me(res http.ResponseWriter,req *http.Request){
	io.WriteString(res,"Welcome Rejul")
}


func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/dog/",dog)
	http.HandleFunc("/me/",me)
	http.ListenAndServe(":8080",nil)
}
