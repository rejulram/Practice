package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	//Request
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			//Headers done
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	//request Line
	m := strings.Fields(ln)[0]
	u := strings.Fields(ln)[1]
	fmt.Println("****Method*****:  ", m)
	fmt.Println("****URL*****:", u)
	if m == "GET" && u == "/" {
		index(conn)
	}
	if m == "GET" && u == "/about" {
		about(conn)
	}
	if m == "GET" && u == "/contact" {
		contact(conn)
	}
	if m == "GET" && u == "/apply" {
		apply(conn)
	}
	if m == "POST" && u == "/apply" {
		applyProcess(conn)
	}
}

func index(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Index</title></head><body>
			<strong>INDEX</strong>
			<a href="/">index</a><br>
			<a href="/about">about</a><br>
			<a href="/contact">contact</a><br>
			<a href="/apply">apply</a><br>
			</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length : %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type : text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>About</title></head><body>
			<strong>ABOUT</strong>
			<a href="/">index</a><br>
			<a href="/about">about</a><br>
			<a href="/contact">contact</a><br>
			<a href="/apply">apply</a><br>
			</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length : %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type : text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contact(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Contact</title></head><body>
			<strong>CONTACT</strong>
			<a href="/">index</a><br>
			<a href="/about">about</a><br>
			<a href="/contact">contact</a><br>
			<a href="/apply">apply</a><br>
			</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length : %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type : text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func apply(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Apply</title></head><body>
			<strong>APPLY</strong>
			<a href="/">index</a><br>
			<a href="/about">about</a><br>
			<a href="/contact">contact</a><br>
			<a href="/apply">apply</a><br>
			<form method ="post" action="/apply">
			<input type="submit" value="apply">
			</form>
			</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length : %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type : text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func applyProcess(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Apply</title></head><body>
			<strong>CONTACT</strong>
			<a href="/">index</a><br>
			<a href="/about">about</a><br>
			<a href="/contact">contact</a><br>
			<a href="/apply">apply</a><br>
			</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length : %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type : text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
