package main

import (
	"io"
	"log"
	"net"
)

func main() {
	lis,err := net.Listen("tcp",":8080")
	if err != nil {
		log.Panic(err)
	}
	defer lis.Close()
	for{
		conn,err := lis.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		io.WriteString(conn,"\n I see you connected \n")
		conn.Close()
	}
}
