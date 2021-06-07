package hexarchsample1project

import (
	"architecture/hexagonal-architecture/hex-arch-sample1-project/internal/service1"
	"context"
)

func Run(ctx context.Context, msg string) {
	if msg == "" {
		msg = "Hello, from external pkg!"
	}

	arg := service1.AppCoreLogicIn{
		From:    "external pkg",
		Message: msg,
	}

	service1.AppCoreLogic(ctx, arg)
}
