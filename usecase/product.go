package usecase

import (
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
		return c.JSON(http.StatusBadRequest, models.GeneralResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid body json",
		})
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
	var (
		product models.Product
		index   int
	)

	id := c.Param("id")

	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, models.GeneralResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid body json",
		})
	}

	for i, p := range products {
		if p.ID == id {
			index = i
		}
	}
	product.ID = id

	if index == 0 {
		return c.JSON(http.StatusNotFound, models.GeneralResponse{
			Status:  http.StatusNotFound,
			Message: "data product not found",
		})
	}

	products[index] = product

	return c.JSON(http.StatusOK, models.GeneralResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    product,
	})
}

func DeleteProduct(c echo.Context) error {
	var index int
	id := c.Param("id")

	for i, p := range products {
		if p.ID == id {
			index = i
		}
	}

	if index < 0 || index > len(products) {
		return c.JSON(http.StatusNotFound, models.GeneralResponse{
			Status:  http.StatusNotFound,
			Message: "data product not found",
		})
	}

	products = append(products[:index], products[index+1:]...)

	return c.JSON(http.StatusOK, models.GeneralResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    id,
	})
}
