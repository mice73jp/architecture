package main

import (
	"architecture/hexagonal-architecture/hex-arch-sample1-project/cmd/http/controller"
	"architecture/hexagonal-architecture/hex-arch-sample1-project/internal/service2"
	"architecture/hexagonal-architecture/hex-arch-sample1-project/internal/service3"
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
	// // MySql
	// c := mysql.Config{
	// 	Host:                 os.Getenv("DB_HSOT"),
	// 	Port:                 os.Getenv("DB_PORT"),
	// 	User:                 os.Getenv("DB_USER"),
	// 	DBName:               os.Getenv("DB_NAME"),
	// 	Passwd:               os.Getenv("DB_PASSWORD"),
	// 	AllowNativePasswords: true,
	// }

	// db, err := mysql.Connect(c)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// Service2
	// gateway2 := service2.NewGateway(db)
	// provider2 := service2.NewProvider(gateway2)
	mockGateway2 := service2.NewMockGateway(service2.NewMockDB())
	provider2 := service2.NewProvider(mockGateway2)

	// Service3
	// gateway3 := service3.NewGateway(db)
	// provider3 := service3.NewProvider(gateway3)
	mockGateway3 := service3.NewMockGateway(service3.NewMockDB())
	provider3 := service3.NewProvider(mockGateway3)

	ctrl := &controller.Controller{}
	ctrl2 := controller.NewController2(provider2)
	ctrl3 := controller.NewController3(provider3)

	e.GET("/:message", ctrl.HandleMessage)
	e.GET("/people/:personID", ctrl2.HandlePersonGet)
	e.POST("/people", ctrl2.HandlePersonRegister)
	e.POST("/accounts", ctrl3.HandleAccountOpen)
	e.POST("/accounts/transfer", ctrl3.HandleMoneyTransfer)
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}
