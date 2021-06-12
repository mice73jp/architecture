package main

import (
	"architecture/hexagonal-architecture/hex-arch-sample1-project/cmd/http/controller"
	"architecture/hexagonal-architecture/hex-arch-sample1-project/internal/service2"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = createMux()

func main() {
	http.Handle("/", e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func init() {
	ctrl := &controller.Controller{}

	mockGateway2 := service2.NewMockGateway(service2.NewMockDB())
	provider2 := service2.NewProvider(mockGateway2)
	ctrl2 := controller.NewController2(provider2)

	e.GET("/:message", ctrl.HandleMessage)
	e.GET("/people/:personID", ctrl2.HandlePersonGet)
	e.POST("/people", ctrl2.HandlePersonRegister)
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}
