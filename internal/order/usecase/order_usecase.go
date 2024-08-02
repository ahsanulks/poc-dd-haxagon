package usecase

import (
	"context"
	"poc/internal/order/entity"
)

type OrderUsecase struct {
	addressGetter entity.AddressGetter
	productGetter entity.ProductGetter
	orderCreator  entity.OrderCreator
}

func NewOrderUsecase(
	addressGetter entity.AddressGetter,
	productGetter entity.ProductGetter,
	orderCreator entity.OrderCreator,
) *OrderUsecase {
	return &OrderUsecase{
		addressGetter: addressGetter,
		productGetter: productGetter,
		orderCreator:  orderCreator,
	}
}

func (o *OrderUsecase) CreateOrder(ctx context.Context, request *CreateOrderRequest) (*entity.Order, error) {
	address, err := o.addressGetter.GetUserAddresses(ctx, request.UserID, request.AddressID)
	if err != nil {
		return nil, err
	}

	products, err := o.getProducts(ctx, request.ProductRequest)
	if err != nil {
		return nil, err
	}

	order, err := entity.CreateOrder(request.UserID, products, address)
	if err != nil {
		return nil, err
	}

	err = o.orderCreator.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *OrderUsecase) getProducts(ctx context.Context, productRequest []*ProductRequest) ([]*entity.Product, error) {
	var productIDs []int
	mapProductQuantity := make(map[int]int)
	for _, product := range productRequest {
		productIDs = append(productIDs, product.ID)
		mapProductQuantity[product.ID] = product.Quantity
	}

	products, err := o.productGetter.GetProducts(ctx, productIDs)
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		product.Quantity = mapProductQuantity[product.ID]
	}
	return products, nil
}
