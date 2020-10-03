package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for v := range gen(ctx) {
		fmt.Println(v)
		if v == 5 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 0
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				dst <- n
				n++
			}
		}

	}()
	return dst
}
