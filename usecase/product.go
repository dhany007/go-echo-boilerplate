package usecase

import (
	"net/http"

	"github.com/dhany007/go-echo-boilerplate/models"
	cst "github.com/dhany007/go-echo-boilerplate/models/constant"
	"github.com/dhany007/go-echo-boilerplate/repository"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ProductUsecase struct {
	Repo *repository.ProductsRepository
}

func NewProductUsecase(r *repository.ProductsRepository) *ProductUsecase {
	return &ProductUsecase{r}
}

func (u *ProductUsecase) AddProduct(c echo.Context) error {
	var args models.CreateProductArgument

	if err := c.Bind(&args); err != nil {
		return ResultError(c, cst.ErrInvalidBodyJSON)
	}

	err := validator.New().Struct(&args)
	if err != nil {
		return ResultError(c, err)
	}

	product, err := u.Repo.AddProduct(args)
	if err != nil {
		return ResultError(c, err)
	}

	return ResultWithData(c, http.StatusCreated, product)
}

func (u *ProductUsecase) GetListProduct(c echo.Context) error {
	products, err := u.Repo.GetListProduct()
	if err != nil {
		return ResultError(c, err)
	}

	return ResultWithData(c, http.StatusOK, products)
}

func (u *ProductUsecase) GetProductByID(c echo.Context) error {
	id := c.Param("id")
	if id == cst.EmptyString {
		return ResultError(c, cst.ErrInvalidParameterID)
	}

	product, err := u.Repo.FindProduct(id)
	if err != nil {
		return ResultError(c, err)
	}

	return ResultWithData(c, http.StatusOK, product)
}

func (u *ProductUsecase) UpdateProduct(c echo.Context) error {
	var product models.Product

	id := c.Param("id")
	if id == cst.EmptyString {
		return ResultError(c, cst.ErrInvalidParameterID)
	}

	if err := c.Bind(&product); err != nil {
		return ResultError(c, cst.ErrInvalidBodyJSON)
	}

	product.ID = id

	err := u.Repo.UpdateProduct(product)
	if err != nil {
		return ResultError(c, err)
	}

	return ResultWithData(c, http.StatusOK, product)
}

func (u *ProductUsecase) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	if id == cst.EmptyString {
		return ResultError(c, cst.ErrInvalidParameterID)
	}

	e := u.Repo.DeleteProduct(id)
	if e != nil {
		return ResultError(c, e)
	}

	return ResultWithData(c, http.StatusOK, id)
}
