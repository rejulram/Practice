package main

import (
	"io"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to main page")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to Dog's Page")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome Rejul")
}

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))
	http.ListenAndServe(":8080", nil)
}
