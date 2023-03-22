package main

import (
	"context"
)

func main() {
	// Create a new context with a cancel function.
	ctx, cancel := context.WithCancel(context.Background())

	StartProducer(ctx, cancel)
}
