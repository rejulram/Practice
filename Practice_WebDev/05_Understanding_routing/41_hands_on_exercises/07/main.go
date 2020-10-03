package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go serve(conn)
	}
}
func serve(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	var i int
	var rMethod, rURI string
	for scanner.Scan() {
		ln := scanner.Text()
		//fmt.Println(ln)
		if i == 0 {
			xs := strings.Fields(ln)
			rMethod = xs[0]
			rURI = xs[1]
			fmt.Println("METHOD", rMethod)
			fmt.Println("URI", rURI)
		}
		if ln == "" {
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
		i++
	}
	body := "CHECKOUT RESPONSE BODY PAYLOAD"
	body += "\n"
	body += rMethod
	body += "\n"
	body += rURI
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn, "Content-Type : text/plain\r\n")
	fmt.Fprintf(conn, "Content-Length : %d \r\n", len(body))
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
