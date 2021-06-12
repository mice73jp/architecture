package main

import (
	"context"

	hexarchsample1project "architecture/hexagonal-architecture/hex-arch-sample1-project"
)

func main() {
	ctx := context.Background()
	msg := "Hello, from external world!"

	hexarchsample1project.Run(ctx, msg)
}
