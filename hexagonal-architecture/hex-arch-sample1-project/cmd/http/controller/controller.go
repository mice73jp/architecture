package controller

import (
	"architecture/hexagonal-architecture/hex-arch-sample1-project/internal/service1"
	"context"

	echo "github.com/labstack/echo/v4"
)

// Controller .
type Controller struct{}

// HandleMessage .
func (ctrl *Controller) HandleMessage(c echo.Context) error {
	msg := c.Param("message")
	if msg == "" {
		msg = "Hello, from http!"
	}

	arg := service1.AppCoreLogicIn{
		From:    "http",
		Message: msg,
	}

	service1.AppCoreLogic(context.Background(), arg)
	return nil
}
