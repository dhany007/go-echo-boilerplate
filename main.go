package main

import (
	"github.com/dhany007/go-echo-boilerplate/repository"
	"github.com/dhany007/go-echo-boilerplate/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/ping", usecase.Ping)

	r := repository.NewProductRepository()
	u := usecase.NewProductUsecase(r)

	e.POST("/products", u.AddProduct)
	// e.GET("/products", usecase.GetListProduct)
	// e.GET("/products/:id", usecase.GetProductByID)
	// e.PUT("/products/:id", usecase.UpdateProduct)
	// e.DELETE("/products/:id", usecase.DeleteProduct)

	e.Logger.Fatal(e.Start(":3000"))
}
