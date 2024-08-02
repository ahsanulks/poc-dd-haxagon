package entity

import "context"

type AddressGetter interface {
	GetUserAddresses(ctx context.Context, userID, addressID int) (*Address, error)
}

type ProductGetter interface {
	GetProducts(ctx context.Context, productIDs []int) ([]*Product, error)
}

type OrderCreator interface {
	CreateOrder(ctx context.Context, order *Order) error
}
