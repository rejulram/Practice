package main

import (
	"io"
	"net/http"
)

type hotdog int
type hotcat int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is dog")
}

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is cat")
}

func main() {
	var d hotdog
	var c hotcat
	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat/", c)
	http.ListenAndServe(":8080", mux)
}
