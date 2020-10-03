package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS		:", runtime.GOOS)
	fmt.Println("Num of CPUs:", runtime.NumCPU())
	fmt.Println("No of GOroutines:", runtime.NumGoroutine())
	wg.Add(2)
	go func() {
		fmt.Println("Hello from one")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello from two")
		wg.Done()
	}()
	fmt.Println("OS		:", runtime.GOOS)
	fmt.Println("Num of CPUs:", runtime.NumCPU())
	fmt.Println("No of GOroutines:", runtime.NumGoroutine())
	wg.Wait()
}
