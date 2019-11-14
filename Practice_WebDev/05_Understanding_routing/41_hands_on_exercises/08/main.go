package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li,err := net.Listen("tcp",":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	for {
		conn,err := li.Accept()
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
	for scanner.Scan(){
		ln := scanner.Text()
		if ln == ""{
			fmt.Println("END OF HEADERS")
			break
		}
	}
	body := `<!DOCTYPE html>
				<html lang="en">
					<head>
						<meta charset="UTF-8">
						<title>Welecome TCP server</title>
					</head>
					<body>
						<h1>HOLY COW THIS IS LOW LEVEL</h1>
					</body>
				</html>`
	io.WriteString(conn,"HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn,"Content-Type : text/html\r\n")
	fmt.Fprintf(conn,"Content-Length : %d \r\n",len(body))
	io.WriteString(conn,"\r\n")
	io.WriteString(conn,body)
}