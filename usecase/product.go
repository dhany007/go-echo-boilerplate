package usecase

import (
	"net/http"

	"github.com/dhany007/go-echo-boilerplate/model"
	"github.com/labstack/echo/v4"
)

func Ping(c echo.Context) error {
	response := model.GeneralResponse{
		Status:  http.StatusOK,
		Message: "success",
	}
	return c.JSON(http.StatusOK, response)
}

func AddProduct(c echo.Context) error {
	return nil
}

func GetListProduct(c echo.Context) error {
	return nil
}

func GetProductByID(c echo.Context) error {
	return nil
}

func UpdateProduct(c echo.Context) error {
	return nil
}

func DeleteProduct(c echo.Context) error {
	return nil
}
