package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func ping(c echo.Context) error {
	return c.JSON(http.StatusOK, response{
		Status:  "success",
		Message: "pong",
	})
}

func addProduct(c echo.Context) error {
	return nil
}

func getListProduct(c echo.Context) error {
	return nil
}

func getProductByID(c echo.Context) error {
	return nil
}

func updateProduct(c echo.Context) error {
	return nil
}

func deleteProduct(c echo.Context) error {
	return nil
}

func main() {
	e := echo.New()

	e.GET("/ping", ping)

	e.POST("/products", addProduct)
	e.GET("/products", getListProduct)
	e.GET("/product/:id", getProductByID)
	e.PUT("/product/:id", updateProduct)
	e.DELETE("/product/:id", deleteProduct)

	e.Logger.Fatal(e.Start(":3000"))
}
