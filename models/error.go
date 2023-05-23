package models

import "errors"

var (
	ErrDataNotFound = errors.New("data not found")
	ErrDataExists   = errors.New("data exists")
)