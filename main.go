package main

import (
	"context"
	"fmt"
)

func doSmt(ctx context.Context) {
	fmt.Println("do something...")
	fmt.Println(ctx.Value("id"))
}
func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "id", "1999")
	doSmt(ctx)
}
