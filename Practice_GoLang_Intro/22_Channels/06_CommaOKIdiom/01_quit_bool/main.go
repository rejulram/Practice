package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)
	//send
	go send(even,odd,quit)
	recieve(even,odd,quit)
}

func send(e,o chan<- int,q chan<- bool){
	for i:=0;i<10;i++{
		if i%2 == 0{
			e <- i
		}else {
			o <- i
		}
	}
	close(q)
}

func recieve(e,o <-chan int, q <-chan bool){
	for {
		select {
		case v:= <-e :
			fmt.Println("From EVEN channel",v)
		case v:=<-o :
			fmt.Println("From ODD channel",v)
		case v,ok:= <-q :
			if !ok {
				fmt.Println("From comma OK bit",v,ok)
				return
			}else{
				fmt.Println("From Comma OK bit")
			}
		}
	}
}