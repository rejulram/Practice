package main

import (
	"fmt"
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
func serve(conn net.Conn){
	defer conn.Close()
}
