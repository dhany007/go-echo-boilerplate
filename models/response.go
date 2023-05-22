package models

type (
	GeneralResponse struct {
		Status  int         `json:"status"`
		Message string      `json:"messages"`
		Data    interface{} `json:"data"`
	}
)
