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
		conn, err := li.Accept()
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
	for scanner.Scan() {
		if text := scanner.Text(); text == ""{
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break;
		}else {
			fmt.Println(text)
		}
	}
	//Code got here
	fmt.Println("Code got here")
	io.WriteString(conn,"\n I see you connected \n")
}

