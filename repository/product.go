package repository

import (
	"github.com/dhany007/go-echo-boilerplate/models"
	cst "github.com/dhany007/go-echo-boilerplate/models/constant"
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
		return res, cst.ErrDataExists
	}

	Products[p.ID] = p

	return p, nil
}

func (r *ProductsRepository) FindProduct(id string) (res models.Product, err error) {
	p, ok := Products[id]
	if !ok {
		return res, cst.ErrDataNotFound
	}

	return p, nil
}

func (r *ProductsRepository) GetListProduct() (res []models.Product, err error) {
	for _, p := range Products {
		res = append(res, p)
	}

	return
}

func (r *ProductsRepository) UpdateProduct(p models.Product) (err error) {
	product, err := r.FindProduct(p.ID)
	if err != nil {
		return err
	}

	Products[p.ID] = product

	return
}

func (r *ProductsRepository) DeleteProduct(id string) (err error) {
	_, e := r.FindProduct(id)
	if e != nil {
		return cst.ErrDataNotFound
	}

	delete(Products, id)

	return
}
