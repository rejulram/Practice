package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	li,err := net.Listen("tcp",":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()
	for{
		conn,err := li.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn){
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		s := strings.ToLower(scanner.Text())
		bs := []byte(s)
		r := rot13(bs)
		fmt.Fprintf(conn,"Rotated value : %s-%s \n \n",s,r)
	}
}

func rot13(bs []byte)[]byte {
	var r13 = make([]byte,len(bs))
	for i,v := range bs {
		//Ascii 97- 122
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13
}