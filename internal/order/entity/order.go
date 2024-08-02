package entity

import "time"

type Order struct {
	ID         int
	UserID     int
	TotalPrice float64
	CreatedAt  time.Time
	Product    []*Product
	Address    *Address
}

type Product struct {
	ID       int
	Quantity int
	Price    float64
}

type Address struct {
	ID        int
	Street    string
	City      string
	ZipCode   string
	Latitude  float64
	Longitude float64
}

func CreateOrder(userID int, products []*Product, address *Address) (*Order, error) {
	if len(products) == 0 {
		return nil, ErrProductRequired
	}

	if address == nil {
		return nil, ErrAddressRequired
	}

	var totalPrice float64
	for _, product := range products {
		totalPrice += float64(product.Quantity) * product.Price
	}

	return &Order{
		UserID:     userID,
		TotalPrice: totalPrice,
		CreatedAt:  time.Time{},
		Product:    products,
		Address:    address,
	}, nil
}
