package main

import (
	"context"
	"fmt"
	"time"
	"runtime"
)

func main() {
	ctx,cancel := context.WithCancel(context.Background())

	fmt.Println("Error check 1 :",ctx.Err())
	fmt.Println("No of GoRoutines:",runtime.NumGoroutine())

	go func(){
		n :=0
		for {
			select {
				case <-ctx.Done():
					return
				default: n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Working",n)
			}
		}
	}()
	time.Sleep(time.Second *2)
	fmt.Println("Error check 2 :",ctx.Err())
	fmt.Println("No of GoRoutines",runtime.NumGoroutine())
    fmt.Println("About to Cancel")
	cancel()
	fmt.Println("Cancelled Context")
	time.Sleep(time.Second)
	fmt.Println("Error check 3 :",ctx.Err())
	fmt.Println("No of GoRoutines",runtime.NumGoroutine())

}
