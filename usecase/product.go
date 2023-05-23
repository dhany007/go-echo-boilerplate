package usecase

import (
	"errors"
	"net/http"

	"github.com/dhany007/go-echo-boilerplate/models"
	"github.com/dhany007/go-echo-boilerplate/repository"
	"github.com/labstack/echo/v4"
)

func Ping(c echo.Context) error {
	response := models.GeneralResponse{
		Status:  http.StatusOK,
		Message: "success",
	}

	return c.JSON(http.StatusOK, response)
}

type ProductUsecase struct {
	Repo *repository.ProductsRepository
}

func NewProductUsecase(r *repository.ProductsRepository) *ProductUsecase {
	return &ProductUsecase{r}
}

func (p *ProductUsecase) AddProduct(c echo.Context) error {
	var args models.Product

	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusBadRequest, models.GeneralResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid body json",
		})
	}

	product, err := p.Repo.AddProduct(args)
	if err != nil {
		e := models.GeneralResponse{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		}

		if errors.Is(err, models.ErrDataExists) {
			e.Status = http.StatusNotFound
			e.Message = "data exists"
		}

		return c.JSON(e.Status, e)
	}

	response := models.GeneralResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data:    product,
	}

	return c.JSON(http.StatusCreated, response)
}

// func GetListProduct(c echo.Context) error {
// 	response := models.GeneralResponse{
// 		Status:  http.StatusOK,
// 		Message: "success",
// 		Data:    products,
// 	}

// 	return c.JSON(http.StatusOK, response)
// }

// func GetProductByID(c echo.Context) error {
// 	var product models.Product

// 	id := c.Param("id")

// 	for _, p := range products {
// 		if p.ID == id {
// 			product = p
// 		}
// 	}

// 	if product.ID == "" {
// 		return c.JSON(http.StatusNotFound, models.GeneralResponse{
// 			Status:  http.StatusNotFound,
// 			Message: "data product not found",
// 		})
// 	}

// 	return c.JSON(http.StatusOK, models.GeneralResponse{
// 		Status:  http.StatusOK,
// 		Message: "success",
// 		Data:    product,
// 	})
// }

// func UpdateProduct(c echo.Context) error {
// 	var (
// 		product models.Product
// 		index   int
// 	)

// 	id := c.Param("id")

// 	if err := c.Bind(&product); err != nil {
// 		return c.JSON(http.StatusBadRequest, models.GeneralResponse{
// 			Status:  http.StatusBadRequest,
// 			Message: "invalid body json",
// 		})
// 	}

// 	for i, p := range products {
// 		if p.ID == id {
// 			index = i
// 		}
// 	}
// 	product.ID = id

// 	if index == 0 {
// 		return c.JSON(http.StatusNotFound, models.GeneralResponse{
// 			Status:  http.StatusNotFound,
// 			Message: "data product not found",
// 		})
// 	}

// 	products[index] = product

// 	return c.JSON(http.StatusOK, models.GeneralResponse{
// 		Status:  http.StatusOK,
// 		Message: "success",
// 		Data:    product,
// 	})
// }

// func DeleteProduct(c echo.Context) error {
// 	var index int
// 	id := c.Param("id")

// 	for i, p := range products {
// 		if p.ID == id {
// 			index = i
// 		}
// 	}

// 	if index < 0 || index > len(products) {
// 		return c.JSON(http.StatusNotFound, models.GeneralResponse{
// 			Status:  http.StatusNotFound,
// 			Message: "data product not found",
// 		})
// 	}

// 	products = append(products[:index], products[index+1:]...)

// 	return c.JSON(http.StatusOK, models.GeneralResponse{
// 		Status:  http.StatusOK,
// 		Message: "success",
// 		Data:    id,
// 	})
// }
