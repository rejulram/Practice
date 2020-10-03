package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Cannot Listen on Port 8080", err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println("Error while accepting the connection", err)
			continue
		}
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			if text := scanner.Text(); text == "" {
				fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
				break
			} else {
				fmt.Println(text)
			}
		}
		//Code got here
		fmt.Println("Code got here")
		conn.Close()
	}
}
