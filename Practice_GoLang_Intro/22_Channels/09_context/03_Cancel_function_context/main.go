package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Println("Context\t:", ctx)
	fmt.Println("Context Err :", ctx.Err())
	fmt.Printf("Context Type :%T\n", ctx)
	fmt.Println("-------------------------")

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("Context\t", ctx)
	fmt.Println("Context Err :", ctx.Err())
	fmt.Printf("Context Type :%T\n", ctx)
	fmt.Println("Cancel\t:", cancel)
	fmt.Printf("Cancel Type  :%T\n", cancel)
	fmt.Println("--------------------------------")
	cancel()

	fmt.Println("Context\t", ctx)
	fmt.Println("Context Err :", ctx.Err())
	fmt.Printf("Context Type :%T\n", ctx)
	fmt.Println("Cancel\t:", cancel)
	fmt.Printf("Cancel Type  :%T\n", cancel)

}
