package main

import (
	"github.com/dhany007/go-echo-boilerplate/repository"
	"github.com/dhany007/go-echo-boilerplate/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	r := repository.NewProductRepository()
	u := usecase.NewProductUsecase(r)

	e.POST("/products", u.AddProduct)
	e.GET("/products", u.GetListProduct)
	e.GET("/products/:id", u.GetProductByID)
	e.PUT("/products/:id", u.UpdateProduct)
	e.DELETE("/products/:id", u.DeleteProduct)

	e.Logger.Fatal(e.Start(":3000"))
}
