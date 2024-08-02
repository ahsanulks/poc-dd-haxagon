package repository

import (
	"context"
	"poc/internal/order/entity"
)

type ProductInMemory struct{}

func NewProductInMemory() *ProductInMemory {
	return &ProductInMemory{}
}

func (r *ProductInMemory) GetProducts(ctx context.Context, productIDs []int) ([]*entity.Product, error) {
	var products []*entity.Product
	for _, id := range productIDs {
		products = append(products, &entity.Product{
			ID:    id,
			Price: 1000,
		})
	}
	return products, nil
}
