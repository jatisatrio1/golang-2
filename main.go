package main

import (
	"rest-api/config"
	"rest-api/controller"

	"github.com/labstack/echo"
)

func main() {
	config.Connect()
	e := echo.New()

	e.GET("/order", controller.ReadOrder)
	e.POST("/order", controller.CreateOrder)
	e.PUT("/order/:id", controller.UpdateOrder)
	e.DELETE("order/:id", controller.DeleteOrder)

	e.Logger.Fatal(e.Start(":1323"))
}
