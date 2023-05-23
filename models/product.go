package models

type (
	Product struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Quantity int64  `json:"quantity"`
		Price    int64  `json:"price"`
	}

	CreateProductArgument struct {
		ID       string `json:"id" validate:"required"`
		Name     string `json:"name" validate:"required,min=3"`
		Quantity int64  `json:"quantity" validate:"required,min=1"`
		Price    int64  `json:"price" validate:"required,min=1"`
	}

	UpdateProductArgument struct {
		Name     string `json:"name"`
		Quantity int64  `json:"quantity"`
		Price    int64  `json:"price"`
	}
)
