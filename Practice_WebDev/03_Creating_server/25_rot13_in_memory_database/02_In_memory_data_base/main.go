package main

import (
	"bufio"
	"fmt"
	"io"
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
	io.WriteString(conn, "\n IN MEMORY DATABASE \n\n"+
		"USE:\n"+
		"SET key value\n"+
		"GET key\n"+
		"DELETE key\n"+
		"Example :\n"+
		"SET fav choclate\n"+
		"GET fav\n")

	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		switch fs[0] {
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "Expecte Value")
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintln(conn, v)
		case "DELETE":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintln(conn, "INVALID COMMAND")
		}
	}

}
