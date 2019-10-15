package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	fmt.Println("CPUs :",runtime.NumCPU())
	fmt.Println("Go Routines:",runtime.NumGoroutine())
	wg.Add(gs)
	for i:=0 ; i< gs ;i++ {
		go func(){
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Go Routines:",runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Go Routines:",runtime.NumGoroutine())
	fmt.Println("Counter :", counter)
}
