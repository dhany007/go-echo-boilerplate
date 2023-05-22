package models

type (
	Product struct {
		ID       string `json:"id,omitempty"`
		Name     string `json:"name"`
		Quantity int64  `json:"quantity"`
		Price    int64  `json:"price"`
	}
)
