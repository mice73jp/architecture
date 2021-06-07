package main

import (
	hexarchsample1project "architecture/hexagonal-architecture/hex-arch-sample1-project"
	"context"
)

func main() {
	ctx := context.Background()
	msg := "Hello, from external world!"

	hexarchsample1project.Run(ctx, msg)
}
