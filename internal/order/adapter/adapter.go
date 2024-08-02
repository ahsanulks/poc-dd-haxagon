package adapter

import (
	"context"
	"poc/internal/order/entity"
	"poc/internal/order/usecase"
)

type OrderUsecase interface {
	CreateOrder(ctx context.Context, request *usecase.CreateOrderRequest) (*entity.Order, error)
}
