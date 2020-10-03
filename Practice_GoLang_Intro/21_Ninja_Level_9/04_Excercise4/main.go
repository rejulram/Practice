package main

import (
	"fmt"
	"runtime"
	"sync"
)

var incrementer int = 0

func main() {
	const gs = 100
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(gs)
	fmt.Println("Num CPUs :", runtime.NumCPU())
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Go Routines :", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("counter :", incrementer)
}
