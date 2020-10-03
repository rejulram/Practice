package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Rejul-Key", "This is Rejul")
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintln(w, "<p>Anything you want in this function</p>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
