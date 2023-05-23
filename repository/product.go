package repository

import (
	"github.com/dhany007/go-echo-boilerplate/models"
)

var (
	Products = make(map[string]models.Product)
)

type ProductsRepository struct{}

func NewProductRepository() *ProductsRepository {
	return &ProductsRepository{}
}

func (r *ProductsRepository) AddProduct(p models.Product) (res models.Product, err error) {
	_, err = r.FindProduct(p.ID)
	if err == nil {
		return res, models.ErrDataExists
	}

	Products[p.ID] = p

	return p, nil
}

func (r *ProductsRepository) FindProduct(id string) (res models.Product, err error) {
	p, ok := Products[id]
	if !ok {
		return res, models.ErrDataNotFound
	}

	return p, nil
}

func (r *ProductsRepository) GetListProduct() (res []models.Product, err error) {
	for _, p := range Products {
		res = append(res, p)
	}

	return
}
