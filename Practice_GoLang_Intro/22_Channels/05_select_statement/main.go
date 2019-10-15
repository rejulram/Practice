package main

import "fmt"

func main() {
	odd := make(chan int)
	even := make(chan int)
	quit := make(chan int)
	// send
	go send(odd,even,quit)
	receive(odd,even,quit)
	//fmt.Println("About to exit")
}

func send(o,e,q chan<- int) {
	for i:=0;i<100;i++ {
		if i%2 == 0 {
			e <- i
		}else {
			o <- i
		}
	}
	//close(e)
	//close(o)
	q <- 1
}

func receive(o,e,q <-chan int){
	for {
		select {
		case v:= <-e :
			fmt.Println("From even channel",v)
		case v:= <-o :
			fmt.Println("From odd channel",v)
		case v:=<-q :
			fmt.Println("From quit channel",v)
			return
		}
	}
}
