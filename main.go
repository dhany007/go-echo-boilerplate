package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response{
			Status:  "success",
			Message: "pong",
		})
	})

	e.Logger.Fatal(e.Start(":3000"))
}
