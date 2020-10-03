package main

import (
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "It is a dog")
	case "/cat":
		io.WriteString(res, "It is a cat")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
