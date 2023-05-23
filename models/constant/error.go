package constant

import "errors"

var (
	DataNotFoundMsg        string = "data not found"
	DataExists             string = "data exists"
	InvalidBodyJsonMsg     string = "invalid body json"
	InternalServerErrorMsg string = "internal server error"
	InvalidParameterIDMsg  string = "invalid parameter id"
)

var (
	ErrDataExists   error = errors.New(DataExists)
	ErrDataNotFound error = errors.New(DataNotFoundMsg)
)
