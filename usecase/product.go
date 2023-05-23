package usecase

import (
	"errors"
	"net/http"

	"github.com/dhany007/go-echo-boilerplate/models"
	cst "github.com/dhany007/go-echo-boilerplate/models/constant"
	"github.com/dhany007/go-echo-boilerplate/repository"
	"github.com/labstack/echo/v4"
)

type ProductUsecase struct {
	Repo *repository.ProductsRepository
}

func NewProductUsecase(r *repository.ProductsRepository) *ProductUsecase {
	return &ProductUsecase{r}
}

func (u *ProductUsecase) AddProduct(c echo.Context) error {
	var args models.Product

	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusBadRequest, models.GeneralResponse{
			Status:  http.StatusBadRequest,
			Message: cst.InvalidBodyJsonMsg,
		})
	}

	product, err := u.Repo.AddProduct(args)
	if err != nil {
		e := models.GeneralResponse{
			Status:  http.StatusInternalServerError,
			Message: cst.InternalServerErrorMsg,
		}

		if errors.Is(err, cst.ErrDataExists) {
			e.Status = http.StatusBadRequest
			e.Message = cst.ErrDataExists.Error()
		}

		return c.JSON(e.Status, e)
	}

	response := models.GeneralResponse{
		Status:  http.StatusCreated,
		Message: cst.Success,
		Data:    product,
	}

	return c.JSON(http.StatusCreated, response)
}

func (u *ProductUsecase) GetListProduct(c echo.Context) error {
	products, err := u.Repo.GetListProduct()
	if err != nil {
		e := models.GeneralResponse{
			Status:  http.StatusInternalServerError,
			Message: cst.InternalServerErrorMsg,
		}

		return c.JSON(e.Status, err)
	}

	response := models.GeneralResponse{
		Status:  http.StatusOK,
		Message: cst.Success,
		Data:    products,
	}

	return c.JSON(http.StatusOK, response)
}

func (u *ProductUsecase) GetProductByID(c echo.Context) error {
	id := c.Param("id")
	if id == cst.EmptyString {
		return c.JSON(http.StatusBadRequest, models.GeneralResponse{
			Status:  http.StatusBadRequest,
			Message: cst.InvalidParameterIDMsg,
		})
	}

	product, err := u.Repo.FindProduct(id)
	if err != nil {
		e := models.GeneralResponse{
			Status:  http.StatusInternalServerError,
			Message: cst.InternalServerErrorMsg,
		}

		if errors.Is(err, cst.ErrDataNotFound) {
			e.Status = http.StatusNotFound
			e.Message = cst.ErrDataNotFound.Error()
		}

		return c.JSON(http.StatusInternalServerError, e)
	}

	return c.JSON(http.StatusOK, models.GeneralResponse{
		Status:  http.StatusOK,
		Message: cst.Success,
		Data:    product,
	})
}

func (u *ProductUsecase) UpdateProduct(c echo.Context) error {
	var product models.Product

	id := c.Param("id")
	if id == cst.EmptyString {
		return c.JSON(http.StatusBadRequest, models.GeneralResponse{
			Status:  http.StatusBadRequest,
			Message: cst.InvalidParameterIDMsg,
		})
	}

	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, models.GeneralResponse{
			Status:  http.StatusBadRequest,
			Message: cst.InvalidBodyJsonMsg,
		})
	}

	product.ID = id

	err := u.Repo.UpdateProduct(product)
	if err != nil {
		e := models.GeneralResponse{
			Status:  http.StatusInternalServerError,
			Message: cst.InternalServerErrorMsg,
		}

		if errors.Is(err, cst.ErrDataNotFound) {
			e.Status = http.StatusNotFound
			e.Message = cst.DataNotFoundMsg
		}

		return c.JSON(e.Status, e)
	}

	return c.JSON(http.StatusOK, models.GeneralResponse{
		Status:  http.StatusOK,
		Message: cst.Success,
		Data:    product,
	})
}

func (u *ProductUsecase) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	if id == cst.EmptyString {
		return c.JSON(http.StatusBadRequest, models.GeneralResponse{
			Status:  http.StatusBadRequest,
			Message: cst.InvalidParameterIDMsg,
		})
	}

	e := u.Repo.DeleteProduct(id)
	if e != nil {
		resp := models.GeneralResponse{
			Status:  http.StatusInternalServerError,
			Message: cst.InternalServerErrorMsg,
		}

		if errors.Is(e, cst.ErrDataNotFound) {
			resp.Status = http.StatusNotFound
			resp.Message = cst.DataNotFoundMsg
		}

		return c.JSON(resp.Status, resp)
	}

	return c.JSON(http.StatusOK, models.GeneralResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    id,
	})
}
