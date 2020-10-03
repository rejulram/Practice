package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Println("Context\t:", ctx)
	fmt.Println("Context Err :", ctx.Err())
	fmt.Printf("Context Type : %T\n", ctx)
	fmt.Println("-------------------------")
	ctx, _ = context.WithCancel(ctx)
	fmt.Println("Context\t:", ctx)
	fmt.Println("Context Err :", ctx.Err())
	fmt.Printf("Context Type : %T\n", ctx)
	fmt.Println("-------------------------")
}
