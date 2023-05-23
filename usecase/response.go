package usecase

import (
	"errors"
	"net/http"

	"github.com/dhany007/go-echo-boilerplate/models"
	cst "github.com/dhany007/go-echo-boilerplate/models/constant"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ResultError(c echo.Context, err error) error {
	var e models.GeneralResponse

	switch {
	case errors.Is(err, cst.ErrDataNotFound):
		e.Status = http.StatusNotFound
		e.Message = cst.DataNotFoundMsg
	case errors.Is(err, cst.ErrDataExists):
		e.Status = http.StatusBadRequest
		e.Message = cst.DataExistsMsg
	case errors.Is(err, cst.ErrInvalidParameterID):
		e.Status = http.StatusBadRequest
		e.Message = cst.InvalidParameterIDMsg
	case errors.Is(err, cst.ErrInvalidBodyJSON):
		e.Status = http.StatusBadRequest
		e.Message = cst.InvalidBodyJSONMsg
	default:
		e.Status = http.StatusInternalServerError
		e.Message = cst.InternalServerErrorMsg

		if _, ok := err.(validator.ValidationErrors); ok {
			e.Status = http.StatusBadRequest
			e.Message = err.Error()
		}
	}

	return c.JSON(e.Status, e)
}

func Result(c echo.Context, code int) error {
	return ResultWithData(c, code, nil)
}

func ResultWithData(c echo.Context, code int, data interface{}) error {
	r := models.GeneralResponse{
		Status:  code,
		Message: cst.Success,
		Data:    data,
	}

	return c.JSON(code, r)
}
