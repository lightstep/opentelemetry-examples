package main

import (
	"context"
)

func main() {
	// Create a new context with a cancel function.
	ctx, _ := context.WithCancel(context.Background())

	StartConsumer(ctx)
}
