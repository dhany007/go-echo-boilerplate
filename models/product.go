package models

type (
	Product struct {
		ID       string `json:"id,omitempty"`
		Name     string `json:"name"`
		Quantity int    `json:"quantity"`
		Price    int    `json:"price"`
	}
)
