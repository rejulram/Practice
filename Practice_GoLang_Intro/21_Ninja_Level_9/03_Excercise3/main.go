package main

import (
	"runtime"
	"sync"
	"fmt"
)

var incrementer int =0

func main() {
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	fmt.Println("Num CPUs :",runtime.NumCPU())
	for i:=0;i< gs ; i++{
		go func(){
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			wg.Done()
		}()
		fmt.Println("Go Routines :",runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("counter :",incrementer)
}
