package main

import (
	"github.com/dhany007/go-echo-boilerplate/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/ping", usecase.Ping)

	e.POST("/products", usecase.AddProduct)
	e.GET("/products", usecase.GetListProduct)
	e.GET("/product/:id", usecase.GetProductByID)
	e.PUT("/product/:id", usecase.UpdateProduct)
	e.DELETE("/product/:id", usecase.DeleteProduct)

	e.Logger.Fatal(e.Start(":3000"))
}
