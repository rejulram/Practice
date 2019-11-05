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
		log.Panic(err)
	}
	defer li.Close()
	for {
		conn,err := li.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	//We never got here
	fmt.Println("Code got here.")
	io.WriteString(conn,"I see you connected")
}
