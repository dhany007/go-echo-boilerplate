package usecase

import (
	"fmt"
	"net/http"

	"github.com/dhany007/go-echo-boilerplate/models"
	"github.com/labstack/echo/v4"
)

var products []models.Product

func Ping(c echo.Context) error {
	response := models.GeneralResponse{
		Status:  http.StatusOK,
		Message: "success",
	}
	return c.JSON(http.StatusOK, response)
}

func AddProduct(c echo.Context) error {
	var product models.Product

	if err := c.Bind(&product); err != nil {
		return err
	}

	for _, p := range products {
		if p.ID == product.ID {
			return c.JSON(http.StatusBadRequest, models.GeneralResponse{
				Status:  http.StatusBadRequest,
				Message: "id product exists",
			})
		}
	}

	products = append(products, product)

	response := models.GeneralResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data:    product,
	}

	return c.JSON(http.StatusCreated, response)
}

func GetListProduct(c echo.Context) error {
	response := models.GeneralResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    products,
	}

	return c.JSON(http.StatusOK, response)
}

func GetProductByID(c echo.Context) error {
	var product models.Product

	id := c.Param("id")
	fmt.Println("id", id)

	for _, p := range products {
		if p.ID == id {
			product = p
		}
	}

	if product.ID == "" {
		return c.JSON(http.StatusNotFound, models.GeneralResponse{
			Status:  http.StatusNotFound,
			Message: "data product not found",
		})
	}

	return c.JSON(http.StatusOK, models.GeneralResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    product,
	})
}

func UpdateProduct(c echo.Context) error {
	return nil
}

func DeleteProduct(c echo.Context) error {
	return nil
}
