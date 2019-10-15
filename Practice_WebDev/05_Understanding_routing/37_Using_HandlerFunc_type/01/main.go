package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter,req *http.Request){
	io.WriteString(res,"This is dog")
}

func c(res http.ResponseWriter, req *http.Request){
	io.WriteString(res,"This is cat")
}

func main() {
	http.Handle("/dog/",http.HandlerFunc(d))
	http.Handle("/cat/",http.HandlerFunc(c))
	http.ListenAndServe(":8080",nil)
}