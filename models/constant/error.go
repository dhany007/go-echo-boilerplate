package constant

import "errors"

var (
	DataNotFoundMsg        string = "data not found"
	DataExistsMsg          string = "data exists"
	InvalidBodyJSONMsg     string = "invalid body json"
	InternalServerErrorMsg string = "internal server error"
	InvalidParameterIDMsg  string = "invalid parameter id"
)

var (
	ErrDataExists         error = errors.New(DataExistsMsg)
	ErrDataNotFound       error = errors.New(DataNotFoundMsg)
	ErrInvalidParameterID error = errors.New(InvalidParameterIDMsg)
	ErrInvalidBodyJSON    error = errors.New(InvalidBodyJSONMsg)
	ErrInternalServer     error = errors.New(InternalServerErrorMsg)
)
