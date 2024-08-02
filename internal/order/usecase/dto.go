package usecase

type CreateOrderRequest struct {
	UserID         int
	ProductRequest []*ProductRequest
	AddressID      int
}

type ProductRequest struct {
	ID       int
	Quantity int
}
