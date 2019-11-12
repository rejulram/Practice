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
			log.Println(err)
			continue
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		if text := scanner.Text(); text == ""{
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break;
		}else {
			fmt.Println()
		}
	}
	body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
	io.WriteString(conn,"HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn,"Content-Type: text/plain\r\n")
	fmt.Fprintf(conn,"Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn,"\r\n")
	io.WriteString(conn,body)
}
