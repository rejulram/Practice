package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var incrementer int64 = 0

func main() {
	var wg sync.WaitGroup
	gs := 100
	wg.Add(100)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final value :", incrementer)
}
